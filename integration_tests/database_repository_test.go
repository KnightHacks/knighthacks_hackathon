package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_hackathon/repository"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

var integrationTest = flag.Bool("integration", false, "whether to run integration tests")
var databaseUri = flag.String("postgres-uri", "postgresql://postgres:test@localhost:5432/postgres", "postgres uri for running integration tests")

var databaseRepository *repository.DatabaseRepository

type Test[A any, T any] struct {
	name    string
	args    A
	want    T
	wantErr bool
}

func TestMain(t *testing.M) {
	flag.Parse()
	// check if integration testing is disabled
	if *integrationTest == false {
		return
	}

	// connect to database
	var err error
	pool, err := database.ConnectWithRetries(*databaseUri)
	fmt.Printf("connecting to database, pool=%v, err=%v\n", pool, err)
	if err != nil {
		log.Fatalf("unable to connect to database err=%v\n", err)
	}

	databaseRepository = repository.NewDatabaseRepository(pool)
	os.Exit(t.Run())
}

func TestDatabaseRepository_AcceptApplicant(t *testing.T) {
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []Test[args, bool]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.AcceptApplicant(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AcceptApplicant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AcceptApplicant() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_ApplyToHackathon(t *testing.T) {

	type args struct {
		ctx         context.Context
		hackathonID string
		userId      string
		input       model.HackathonApplicationInput
	}
	tests := []Test[args, bool]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.ApplyToHackathon(tt.args.ctx, tt.args.hackathonID, tt.args.userId, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyToHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ApplyToHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_CreateHackathon(t *testing.T) {

	type args struct {
		ctx   context.Context
		input *model.HackathonCreateInput
	}
	tests := []Test[args, *model.Hackathon]{
		{
			name: "Create 2023 Hackathon",
			args: args{
				ctx: context.Background(),
				input: &model.HackathonCreateInput{
					Year: 2023,
					Semester: model.SemesterFall,
					StartDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
					EndDate: time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &model.Hackathon{
				Term: &model.Term{
					Year: 2023,
					Semester: model.SemesterFall,
				},
				StartDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
				EndDate: time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.CreateHackathon(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Term, tt.want.Term) || !reflect.DeepEqual(got.StartDate, tt.want.StartDate) || !reflect.DeepEqual(got.EndDate, tt.want.EndDate){
				t.Errorf("CreateHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DeleteHackathon(t *testing.T) {

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []Test[args, bool]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.DeleteHackathon(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DenyApplicant(t *testing.T) {
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []Test[args, bool]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.DenyApplicant(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DenyApplicant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DenyApplicant() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetApplication(t *testing.T) {
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []Test[args, *model.HackathonApplication]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetApplication(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetApplicationWithQueryable(t *testing.T) {

	type args struct {
		ctx         context.Context
		queryable   database.Queryable
		hackathonID string
		userID      string
	}
	tests := []Test[args, *model.HackathonApplication]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetApplicationWithQueryable(tt.args.ctx, tt.args.queryable, tt.args.hackathonID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationWithQueryable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationWithQueryable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetApplicationsByHackathon(t *testing.T) {
	type args struct {
		ctx    context.Context
		obj    *model.Hackathon
		first  int
		after  *string
		status model.ApplicationStatus
	}
	type want struct {
		applications []*model.HackathonApplication
		total        int
	}
	tests := []Test[args, want]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			applications, total, err := databaseRepository.GetApplicationsByHackathon(tt.args.ctx, tt.args.obj, tt.args.first, tt.args.after, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationsByHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(applications, tt.want.applications) {
				t.Errorf("GetApplicationsByHackathon() applications = %v, want %v", applications, tt.want)
			}
			if total != tt.want.total {
				t.Errorf("GetApplicationsByHackathon() total = %v, want %v", total, tt.want.total)
			}
		})
	}
}

func TestDatabaseRepository_GetApplicationsByUser(t *testing.T) {

	type args struct {
		ctx context.Context
		obj *model.User
	}
	tests := []Test[args, *[]model.HackathonApplication]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetApplicationsByUser(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationsByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetCurrentHackathon(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []Test[args, *model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetCurrentHackathon(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathon(t *testing.T) {

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []Test[args, *model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetHackathon(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonByEvent(t *testing.T) {

	type args struct {
		ctx context.Context
		obj *model.Event
	}
	tests := []Test[args, *model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetHackathonByEvent(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonByEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathonByEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonByTermYearAndTermSemester(t *testing.T) {

	type args struct {
		ctx          context.Context
		termYear     int
		termSemester model.Semester
	}
	tests := []Test[args, *model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetHackathonByTermYearAndTermSemester(tt.args.ctx, tt.args.termYear, tt.args.termSemester)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonByTermYearAndTermSemester() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathonByTermYearAndTermSemester() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonEvents(t *testing.T) {

	type args struct {
		ctx       context.Context
		hackathon *model.Hackathon
		first     int
		after     string
	}
	type want struct {
		events []*model.Event
		total  int
	}

	tests := []Test[args, want]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events, total, err := databaseRepository.GetHackathonEvents(tt.args.ctx, tt.args.hackathon, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(events, tt.want.events) {
				t.Errorf("GetHackathonEvents() sponsors = %v, want %v", events, tt.want)
			}
			if total != tt.want.total {
				t.Errorf("GetHackathonEvents() total = %v, want %v", total, tt.want.total)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonSponsors(t *testing.T) {

	type args struct {
		ctx       context.Context
		hackathon *model.Hackathon
		first     int
		after     string
	}
	type want struct {
		sponsors []*model.Sponsor
		total    int
	}

	tests := []Test[args, want]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			sponsors, total, err := databaseRepository.GetHackathonSponsors(tt.args.ctx, tt.args.hackathon, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonSponsors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(sponsors, tt.want.sponsors) {
				t.Errorf("GetHackathonSponsors() sponsors = %v, want %v", sponsors, tt.want)
			}
			if total != tt.want.total {
				t.Errorf("GetHackathonSponsors() total = %v, want %v", total, tt.want.total)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathons(t *testing.T) {

	type args struct {
		ctx    context.Context
		filter *model.HackathonFilter
	}
	tests := []Test[args, []*model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetHackathons(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathons() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonsBySponsor(t *testing.T) {

	type args struct {
		ctx context.Context
		obj *model.Sponsor
	}
	tests := []Test[args, []*model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetHackathonsBySponsor(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonsBySponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathonsBySponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_UpdateApplicantStatus(t *testing.T) {

	type args struct {
		ctx         context.Context
		queryable   database.Queryable
		hackathonID string
		userID      string
		status      model.ApplicationStatus
	}
	tests := []Test[args, any]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateApplicantStatus(tt.args.ctx, tt.args.queryable, tt.args.hackathonID, tt.args.userID, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateApplicantStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateApplication(t *testing.T) {

	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
		input       model.HackathonApplicationInput
	}
	tests := []Test[args, *model.HackathonApplication]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.UpdateApplication(tt.args.ctx, tt.args.hackathonID, tt.args.userID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_UpdateHackathon(t *testing.T) {

	type args struct {
		ctx   context.Context
		id    string
		input *model.HackathonUpdateInput
	}
	tests := []Test[args, *model.Hackathon]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.UpdateHackathon(tt.args.ctx, tt.args.id, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_getTermById(t *testing.T) {

	type args struct {
		ctx       context.Context
		queryable database.Queryable
		id        int
	}
	tests := []Test[args, *model.Term]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetTermById(tt.args.ctx, tt.args.queryable, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTermById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTermById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_getTermId(t *testing.T) {

	type args struct {
		ctx          context.Context
		queryable    database.Queryable
		termYear     int
		termSemester model.Semester
	}
	tests := []Test[args, int]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetTermId(tt.args.ctx, tt.args.queryable, tt.args.termYear, tt.args.termSemester)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTermId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTermId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDatabaseRepository(t *testing.T) {
	type args struct {
		databasePool *pgxpool.Pool
	}
	tests := []Test[args, *repository.DatabaseRepository]{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewDatabaseRepository(tt.args.databasePool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
