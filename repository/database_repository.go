package repository

import (
	"context"
	"errors"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/structure"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//DatabaseRepository
//Implements the Repository interface's functions
type DatabaseRepository struct {
	DatabasePool *pgxpool.Pool
	TermBiMap    *structure.BiMap
}

var (
	NoHackathonByTerm = errors.New("unable to find hackathon by term")
)

func NewDatabaseRepository(databasePool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePool: databasePool,
		TermBiMap:    structure.NewBiMap(),
	}
}

func (r *DatabaseRepository) CreateHackathon(ctx context.Context, input *model.HackathonCreateInput) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) UpdateHackathon(ctx context.Context, id string, input *model.HackathonUpdateInput) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
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

		_, err = r.getTermId(ctx, queryable, termYear, termSemester)
		if err != nil {
			if errors.Is(err, NoHackathonByTerm) {
				return nil, nil
			}
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
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
		return 0, err
	}
	if termId == nil {
		return 0, NoHackathonByTerm
	}
	return *termId, nil
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

func (r *DatabaseRepository) GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}
