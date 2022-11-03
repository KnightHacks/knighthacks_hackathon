package repository

import (
	"context"
	"errors"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/structure"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"

	"strconv"
	"time"
)

// DatabaseRepository
// Implements the Repository interface's functions
type DatabaseRepository struct {
	DatabasePool *pgxpool.Pool
	TermBiMap    *structure.BiMap
}

var (
	NoHackathonByTerm = errors.New("unable to find hackathon by term")
	HackathonNotFound = errors.New("hackathon not found")
)

func NewDatabaseRepository(databasePool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePool: databasePool,
		TermBiMap:    structure.NewBiMap(),
	}
}

func (r *DatabaseRepository) CreateHackathon(ctx context.Context, input *model.HackathonCreateInput) (*model.Hackathon, error) {
	// TODO: Implement handling of Sponsors & Events, pretty sure these lists will be empty...
	term := model.Term{
		Year:     input.Year,
		Semester: input.Semester,
	}

	var queryable database.Queryable = r.DatabasePool

	termId, ok := r.TermBiMap.Get(term).(int)
	var err error
	if !ok {
		tx, err := r.DatabasePool.Begin(ctx)
		if err != nil {
			return nil, err
		}
		queryable = tx

		termId, err = r.getTermId(ctx, queryable, term.Year, term.Semester)
		if err != nil {
			if errors.Is(err, NoHackathonByTerm) {
				err = queryable.QueryRow(
					ctx,
					"INSERT INTO terms (year, semester) VALUES ($1, $2) RETURNING id",
					input.Year,
					input.Semester.String(),
				).Scan(&termId)
				if err != nil {
					return nil, err
				}
			}
			return nil, err
		}
		r.TermBiMap.Put(termId, term)
	}

	var hackathonIdInt int
	if err = queryable.QueryRow(
		ctx,
		"INSERT INTO hackathons (term_id, start_date, end_date) VALUES ($1, $2, $3) RETURNING id",
		termId,
		input.StartDate,
		input.EndDate,
	).Scan(&hackathonIdInt); err != nil {
		return nil, err
	}

	if tx, ok := queryable.(pgx.Tx); ok {
		if err = tx.Commit(ctx); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	return &model.Hackathon{
		ID:        strconv.Itoa(hackathonIdInt),
		Term:      &term,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}, nil
}

func (r *DatabaseRepository) UpdateHackathon(ctx context.Context, id string, input *model.HackathonUpdateInput) (*model.Hackathon, error) {
	if input.Year == nil &&
		input.Semester == nil &&
		len(input.AddedEvents) == 0 &&
		len(input.RemovedEvents) == 0 &&
		len(input.AddedSponsors) == 0 &&
		len(input.RemovedSponsors) == 0 &&
		len(input.AddedParticipants) == 0 &&
		len(input.RemovedParticipants) == 0 {
		return nil, errors.New("empty input field")
	}
	var hackathon *model.Hackathon
	var err error

	tx, err := r.DatabasePool.Begin(ctx)
	runTx := func(tx pgx.Tx, hackathonIdString string, input *model.HackathonUpdateInput) (err error) {
		hackathonId, err := strconv.Atoi(hackathonIdString)
		if err != nil {
			return err
		}
		if input.Year != nil {
			if err = r.updateHackathonYear(ctx, tx, hackathonId, *input.Year); err != nil {
				return err
			}
		}
		if input.Semester != nil {
			if err = r.updateHackathonSemester(ctx, tx, hackathonId, *input.Semester); err != nil {
				return err
			}
		}

		if len(input.AddedEvents) > 0 {
			if err = r.addHackathonEvents(ctx, tx, hackathonId, input.AddedEvents); err != nil {
				return err
			}
		}
		if len(input.RemovedEvents) > 0 {
			if err = r.removeHackathonEvents(ctx, tx, input.RemovedEvents); err != nil {
				return err
			}
		}

		if len(input.AddedSponsors) > 0 {
			if err = r.addHackathonSponsors(ctx, tx, hackathonId, input.AddedSponsors); err != nil {
				return err
			}
		}
		if len(input.RemovedSponsors) > 0 {
			if err = r.removeHackathonSponsors(ctx, tx, hackathonId, input.RemovedSponsors); err != nil {
				return err
			}
		}

		if len(input.AddedParticipants) > 0 {
			if err = r.addHackathonParticipants(ctx, tx, hackathonId, input.AddedParticipants); err != nil {
				return err
			}
		}
		if len(input.RemovedParticipants) > 0 {
			if err = r.removeHackathonParticipants(ctx, tx, hackathonId, input.RemovedParticipants); err != nil {
				return err
			}
		}
		hackathon, err = r.getHackathon(ctx, r.DatabasePool, "SELECT id, term_id, start_date, end_date FROM hackathons WHERE id = $1", id)

		if err != nil {
			return err
		}
		return nil
	}

	err = runTx(tx, id, input)
	if err != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
	}

	return hackathon, nil
}

func (r *DatabaseRepository) updateHackathonYear(ctx context.Context, tx pgx.Tx, hackathonId int, year int) error {
	// This sql statement updates the semester in the terms table where the id of the term row equals
	// the term_id from the hackathons table where the id of that row is equal to the supplied hackathonId
	exec, err := tx.Exec(
		ctx,
		"UPDATE terms SET year = $1 WHERE id IN (SELECT hackathons.term_id FROM hackathons WHERE hackathons.id = $2)",
		year,
		hackathonId,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return HackathonNotFound
	}
	return nil
}

func (r *DatabaseRepository) updateHackathonSemester(ctx context.Context, tx pgx.Tx, hackathonId int, semester model.Semester) error {
	// This sql statement updates the semester in the terms table where the id of the term row equals
	// the term_id from the hackathons table where the id of that row is equal to the supplied hackathonId
	exec, err := tx.Exec(
		ctx,
		"UPDATE terms SET semester = $1 WHERE id IN (SELECT hackathons.term_id FROM hackathons WHERE hackathons.id = $2)",
		semester.String(),
		hackathonId,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return HackathonNotFound
	}
	return nil
}

func (r *DatabaseRepository) addHackathonEvents(ctx context.Context, tx pgx.Tx, hackathonId int, events []string) error {
	for _, eventId := range events {
		if err := r.updateHackathonEvent(ctx, tx, eventId, &hackathonId); err != nil {
			return err
		}
	}
	return nil
}

func (r *DatabaseRepository) removeHackathonEvents(ctx context.Context, tx pgx.Tx, events []string) error {
	for _, eventId := range events {
		if err := r.updateHackathonEvent(ctx, tx, eventId, nil); err != nil {
			return err
		}
	}
	return nil
}

func (r *DatabaseRepository) updateHackathonEvent(ctx context.Context, tx pgx.Tx, eventId string, value *int) error {
	exec, err := tx.Exec(ctx, "UPDATE events SET hackathon_id = $1 WHERE id = $2", value, eventId)
	if err != nil {
		return err
	}
	if exec.RowsAffected() == 0 {
		return errors.New("unable to find event")
	}
	return nil
}

func (r *DatabaseRepository) addHackathonSponsors(ctx context.Context, tx pgx.Tx, hackathonId int, sponsors []string) error {
	for _, sponsorId := range sponsors {
		_, err := tx.Exec(ctx, "INSERT INTO hackathon_sponsors (hackathon_id, sponsor_id) VALUES ($1, $2)", hackathonId, sponsorId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *DatabaseRepository) removeHackathonSponsors(ctx context.Context, tx pgx.Tx, hackathonId int, sponsors []string) error {
	for _, sponsorId := range sponsors {
		_, err := tx.Exec(ctx, "DELETE FROM hackathon_sponsors WHERE hackathon_id = $1 AND sponsor_id = $2", hackathonId, sponsorId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *DatabaseRepository) addHackathonParticipants(ctx context.Context, tx pgx.Tx, hackathonId int, participants []string) error {
	for _, participantId := range participants {
		_, err := tx.Exec(ctx, "INSERT INTO hackathon_participants (hackathon_id, user_id) VALUES ($1, $2)", hackathonId, participantId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *DatabaseRepository) removeHackathonParticipants(ctx context.Context, tx pgx.Tx, hackathonId int, participants []string) error {
	for _, participantId := range participants {
		_, err := tx.Exec(ctx, "DELETE FROM hackathon_participants WHERE hackathon_id = $1 AND user_id = $2", hackathonId, participantId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *DatabaseRepository) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	return r.getHackathon(ctx, r.DatabasePool, "SELECT id, term_id, start_date, end_date FROM hackathons WHERE id = $1", id)
}

func (r *DatabaseRepository) GetHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error) {
	var queryable database.Queryable = r.DatabasePool

	term := model.Term{
		Year:     termYear,
		Semester: termSemester,
	}
	termId, ok := r.TermBiMap.Get(term).(int)
	if !ok {
		tx, err := r.DatabasePool.Begin(ctx)
		if err != nil {
			return nil, err
		}
		queryable = tx

		termId, err = r.getTermId(ctx, queryable, termYear, termSemester)
		if err != nil {
			if errors.Is(err, NoHackathonByTerm) {
				return nil, nil
			}
			return nil, err
		}
		r.TermBiMap.Put(termId, term)
		defer tx.Commit(ctx).Error()
	}

	hackathon, err := r.getHackathon(ctx, queryable, "SELECT id, start_date, end_date FROM hackathons WHERE term_id = $1", termId)
	if err != nil {
		return nil, err
	}
	hackathon.Term = &term
	return hackathon, nil
}

func (r *DatabaseRepository) getHackathon(ctx context.Context, queryable database.Queryable, sql string, args ...any) (*model.Hackathon, error) {
	var hackathon model.Hackathon
	err := queryable.QueryRow(ctx, sql, args).Scan(
		&hackathon.ID,
		&hackathon.StartDate,
		&hackathon.EndDate,
	)
	if err != nil {
		return nil, err
	}
	return &hackathon, nil
}

func (r *DatabaseRepository) getTermId(ctx context.Context, queryable database.Queryable, termYear int, termSemester model.Semester) (int, error) {
	var termId *int
	err := queryable.QueryRow(ctx, "SELECT id FROM terms WHERE year = $1 AND semester = $2", termYear, termSemester.String()).Scan(termId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, NoHackathonByTerm
		}
	}
	return *termId, nil
}

func (r *DatabaseRepository) getTermById(ctx context.Context, queryable database.Queryable, id int) (*model.Term, error) {
	var term model.Term
	err := queryable.QueryRow(ctx, "SELECT year, semester FROM terms WHERE id = $1", id).Scan(&term.Year, &term.Semester)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return nil, nil
}

func (r *DatabaseRepository) DeleteHackathon(ctx context.Context, id string) (bool, error) {
	exec, err := r.DatabasePool.Exec(ctx, "DELETE FROM hackathons WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	if exec.RowsAffected() != 1 {
		return false, nil
	}
	return true, nil
}

// GetCurrentHackathon
// TODO: Change name to GetNextHackathon
func (r *DatabaseRepository) GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	var hackathon model.Hackathon
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var termId int
		// TODO: Check validity of using DESC
		err := tx.QueryRow(ctx, "SELECT id, term_id, start_date, end_date FROM hackathons WHERE end_date > CURRENT_DATE ORDER BY end_date DESC LIMIT 1").Scan(
			&hackathon.ID,
			&termId,
			&hackathon.StartDate,
			&hackathon.EndDate,
		)
		if err != nil {
			return err
		}
		term, ok := r.TermBiMap.Get(termId).(model.Term)
		if ok {
			hackathon.Term = &term
		} else {
			term, err := r.getTermById(ctx, tx, termId)
			if err != nil {
				return err
			}
			r.TermBiMap.Put(termId, term)
			hackathon.Term = term
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &hackathon, nil
}

func (r *DatabaseRepository) GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error) {
	var rows pgx.Rows
	var err error

	if filter.Semester != nil {
		query := `
SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         INNER JOIN terms ON hackathons.term_id = terms.id
WHERE terms.year = $1
  AND terms.semester = $2`
		rows, err = r.DatabasePool.Query(ctx, query, filter.Year, filter.Semester)
	} else {
		query := `
SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         INNER JOIN terms ON hackathons.term_id = terms.id
WHERE terms.year = $1`
		rows, err = r.DatabasePool.Query(ctx, query, filter.Year)
	}
	if err != nil {
		return nil, err
	}

	hackathons := make([]*model.Hackathon, 0, 10)

	for rows.Next() {
		var hackathon = model.Hackathon{Term: new(model.Term)}
		var termId int
		err = rows.Scan(
			&hackathon.ID,
			&hackathon.StartDate,
			&hackathon.EndDate,
			&termId,
			&hackathon.Term.Semester,
			&hackathon.Term.Year,
		)
		if err != nil {
			return nil, err
		}
		r.TermBiMap.Put(termId, hackathon.Term)
		hackathons = append(hackathons, &hackathon)
	}

	return hackathons, err
}

func (r *DatabaseRepository) AcceptApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var acceptedDate *time.Time
		err := tx.QueryRow(ctx, "SELECT accepted_date FROM hackathon_participants WHERE user_id = $1 AND hackathon_id = $2", hackathonID, userID).Scan(acceptedDate)
		if err != nil {
			return err
		}
		log.Printf("acceptedDate=%v\n", acceptedDate)
		if acceptedDate == nil {
			// TODO: check how this is handled, maybe the accepted is null in postgres but the pgx representation is just 0 time?
			_, err = tx.Exec(ctx, "INSERT INTO hackathon_participants (user_id, hackathon_id, accepted_date) VALUES ($1, $2, $3)", userID, hackathonID, time.Now().UTC())
			if err != nil {
				return err
			}
		} else {
			return errors.New("this user has already been accepted to this hackathon")
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *DatabaseRepository) DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var acceptedDate *time.Time
		err := tx.QueryRow(ctx, "SELECT accepted_date FROM hackathon_participants WHERE user_id = $1 AND hackathon_id = $2", hackathonID, userID).Scan(acceptedDate)
		if err != nil {
			return err
		}
		log.Printf("acceptedDate=%v\n", acceptedDate)
		if acceptedDate != nil {
			// TODO: check how this is handled, maybe the accepted is null in postgres but the pgx representation is just 0 time?
			return errors.New("this user has already been accepted to this hackathon")
		} else {
			_, err = tx.Exec(ctx, "DELETE FROM hackathon_participants WHERE user_id = $1 AND hackathon_id = $2", userID, hackathonID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *DatabaseRepository) GetHackathonsBySponsor(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error) {
	query := `
SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         FULL JOIN terms ON hackathons.term_id = terms.id
         INNER JOIN hackathon_sponsors on hackathons.id = hackathon_sponsors.hackathon_id
WHERE hackathon_sponsors.sponsor_id = $1`

	intId, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}

	rows, err := r.DatabasePool.Query(ctx, query, intId)
	hackathons := make([]*model.Hackathon, 0, 10)

	for rows.Next() {
		var hackathon = model.Hackathon{Term: new(model.Term)}
		var termId int
		err = rows.Scan(
			&hackathon.ID,
			&hackathon.StartDate,
			&hackathon.EndDate,
			&termId,
			&hackathon.Term.Semester,
			&hackathon.Term.Year,
		)
		if err != nil {
			return nil, err
		}
		r.TermBiMap.Put(termId, hackathon.Term)
		hackathons = append(hackathons, &hackathon)
	}

	return hackathons, err
}

func (r *DatabaseRepository) GetHackathonsByUser(ctx context.Context, obj *model.User, attended bool) ([]*model.Hackathon, error) {
	var query string

	if attended {
		query = `SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         FULL JOIN terms ON hackathons.term_id = terms.id
         INNER JOIN hackathon_participants on hackathons.id = hackathon_participants.hackathon_id
WHERE hackathon_participants.user_id = $1 AND hackathon_participants.accepted_date IS NOT NULL`
	} else {
		query = `SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         FULL JOIN terms ON hackathons.term_id = terms.id
         INNER JOIN hackathon_participants on hackathons.id = hackathon_participants.hackathon_id
WHERE hackathon_participants.user_id = $1 AND hackathon_participants.accepted_date IS NULL`
	}

	intId, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}

	rows, err := r.DatabasePool.Query(ctx, query, intId)
	hackathons := make([]*model.Hackathon, 0, 10)

	for rows.Next() {
		var hackathon = model.Hackathon{Term: new(model.Term)}
		var termId int
		err = rows.Scan(
			&hackathon.ID,
			&hackathon.StartDate,
			&hackathon.EndDate,
			&termId,
			&hackathon.Term.Semester,
			&hackathon.Term.Year,
		)
		if err != nil {
			return nil, err
		}
		r.TermBiMap.Put(termId, hackathon.Term)
		hackathons = append(hackathons, &hackathon)
	}

	return hackathons, err
}

func (r *DatabaseRepository) GetHackathonByEvent(ctx context.Context, obj *model.Event) (*model.Hackathon, error) {
	intId, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}
	var hackathon = model.Hackathon{Term: new(model.Term)}
	var termId int
	row := r.DatabasePool.QueryRow(ctx, `SELECT hackathons.id,
       hackathons.start_date,
       hackathons.end_date,
       terms.id,
       terms.semester,
       terms.year
FROM hackathons
         FULL JOIN terms ON hackathons.term_id = terms.id
         INNER JOIN events on hackathons.id = events.hackathon_id
WHERE events.id = $1`, intId)

	err = row.Scan(&hackathon.ID,
		&hackathon.StartDate,
		&hackathon.EndDate,
		&termId,
		&hackathon.Term.Semester,
		&hackathon.Term.Year)
	if err != nil {
		return nil, err
	}
	r.TermBiMap.Put(termId, hackathon.Term)
	return &hackathon, err
}

// TODO: Change attending/pending to a user status enum

func (r *DatabaseRepository) IsUserAttending(ctx context.Context, hackathon *model.Hackathon, userID string) (bool, error) {
	var exists int
	err := r.DatabasePool.QueryRow(ctx, `SELECT 1
FROM hackathon_participants
WHERE user_id = $1
  AND hackathon_id = $2
  AND accepted_date IS NOT NULL`, userID, hackathon.ID).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (r *DatabaseRepository) IsUserPending(ctx context.Context, hackathon *model.Hackathon, userID string) (bool, error) {
	var exists int
	err := r.DatabasePool.QueryRow(ctx, `SELECT 1
FROM hackathon_participants
WHERE user_id = $1
  AND hackathon_id = $2
  AND accepted_date IS NULL`, userID, hackathon.ID).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (r *DatabaseRepository) GetHackathonApplicants(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.User, int, error) {
	return r.getHackathonUsers(ctx, hackathon, first, after, false)
}

func (r *DatabaseRepository) GetHackathonAttendees(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.User, int, error) {
	return r.getHackathonUsers(ctx, hackathon, first, after, true)
}

func (r *DatabaseRepository) getHackathonUsers(ctx context.Context, hackathon *model.Hackathon, first int, after string, attending bool) ([]*model.User, int, error) {
	var countQuery string
	var userQuery string
	if attending {
		countQuery = `SELECT COUNT(*) FROM hackathon_participants WHERE hackathon_id = $1 AND accepted_date IS NOT NULL`
		userQuery = `SELECT user_id FROM hackathon_participants WHERE hackathon_id = $1 AND accepted_date IS NOT NULL AND user_id > $2 ORDER BY user_id DESC LIMIT $3`
	} else {
		countQuery = `SELECT COUNT(*) FROM hackathon_participants WHERE hackathon_id = $1 AND accepted_date IS NULL`
		userQuery = `SELECT user_id FROM hackathon_participants WHERE hackathon_id = $1 AND accepted_date IS NULL AND user_id > $2 ORDER BY user_id DESC LIMIT $3`
	}

	users := make([]*model.User, 0, first)
	var total int
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(
			ctx,
			userQuery,
			hackathon.ID,
			after,
			first,
		)
		if err != nil {
			return err
		}
		for rows.Next() {
			var user model.User
			err = rows.Scan(&user.ID)
			if err != nil {
				return err
			}
			users = append(users, &user)
		}
		if err = rows.Err(); err != nil {
			return err
		}
		err = tx.QueryRow(ctx, countQuery, hackathon.ID).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return users, total, err
}

func (r *DatabaseRepository) GetHackathonSponsors(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Sponsor, int, error) {
	sponsors := make([]*model.Sponsor, 0, first)
	var total int
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(
			ctx,
			`SELECT sponsor_id FROM hackathon_sponsors WHERE hackathon_id = $1 AND sponsor_id > $2 ORDER BY sponsor_id DESC LIMIT $3`,
			hackathon.ID,
			after,
			first,
		)
		if err != nil {
			return err
		}
		for rows.Next() {
			var sponsor model.Sponsor
			err = rows.Scan(&sponsor.ID)
			if err != nil {
				return err
			}
			sponsors = append(sponsors, &sponsor)
		}
		if err = rows.Err(); err != nil {
			return err
		}
		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM hackathon_sponsors WHERE hackathon_id = $1`, hackathon.ID).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return sponsors, total, err
}

func (r *DatabaseRepository) GetHackathonEvents(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Event, int, error) {
	events := make([]*model.Event, 0, first)
	var total int
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(
			ctx,
			`SELECT id FROM events WHERE hackathon_id = $1 AND id > $2 ORDER BY id DESC LIMIT $3`,
			hackathon.ID,
			after,
			first,
		)
		if err != nil {
			return err
		}
		for rows.Next() {
			var event model.Event
			err = rows.Scan(&event.ID)
			if err != nil {
				return err
			}
			events = append(events, &event)
		}
		if err = rows.Err(); err != nil {
			return err
		}
		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM events WHERE hackathon_id = $1`, hackathon.ID).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return events, total, err
}

func (r *DatabaseRepository) GetApplicationsByUser(ctx context.Context, obj *model.User) ([]*model.HackathonApplication, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetApplication(ctx context.Context, hackathonID string, userID string) (*model.HackathonApplication, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) ApplyToHackathon(ctx context.Context, hackathonID string, input model.HackathonApplicationInput) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) UpdateApplication(ctx context.Context, hackathonID string, userID string, input model.HackathonApplicationInput) (*model.HackathonApplication, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetApplicationsByHackathon(ctx context.Context, obj *model.Hackathon, first int, after *string, status model.ApplicationStatus) ([]*model.HackathonApplication, int, error) {
	//TODO implement me
	panic("implement me")
}
