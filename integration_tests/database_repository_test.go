package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_hackathon/repository"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/structure"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"reflect"
	"testing"
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.AcceptApplicant(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		hackathonID string
		userId      string
		input       model.HackathonApplicationInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.ApplyToHackathon(tt.args.ctx, tt.args.hackathonID, tt.args.userId, tt.args.input)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx   context.Context
		input *model.HackathonCreateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.CreateHackathon(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHackathon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DeleteHackathon(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.DeleteHackathon(tt.args.ctx, tt.args.id)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.DenyApplicant(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.HackathonApplication
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetApplication(tt.args.ctx, tt.args.hackathonID, tt.args.userID)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		queryable   database.Queryable
		hackathonID string
		userID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.HackathonApplication
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetApplicationWithQueryable(tt.args.ctx, tt.args.queryable, tt.args.hackathonID, tt.args.userID)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx    context.Context
		obj    *model.Hackathon
		first  int
		after  *string
		status model.ApplicationStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.HackathonApplication
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, got1, err := r.GetApplicationsByHackathon(tt.args.ctx, tt.args.obj, tt.args.first, tt.args.after, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationsByHackathon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationsByHackathon() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetApplicationsByHackathon() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseRepository_GetApplicationsByUser(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
		obj *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.HackathonApplication
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetApplicationsByUser(tt.args.ctx, tt.args.obj)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetCurrentHackathon(tt.args.ctx)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetHackathon(tt.args.ctx, tt.args.id)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
		obj *model.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetHackathonByEvent(tt.args.ctx, tt.args.obj)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx          context.Context
		termYear     int
		termSemester model.Semester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetHackathonByTermYearAndTermSemester(tt.args.ctx, tt.args.termYear, tt.args.termSemester)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx       context.Context
		hackathon *model.Hackathon
		first     int
		after     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Event
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, got1, err := r.GetHackathonEvents(tt.args.ctx, tt.args.hackathon, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathonEvents() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetHackathonEvents() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathonSponsors(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx       context.Context
		hackathon *model.Hackathon
		first     int
		after     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Sponsor
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, got1, err := r.GetHackathonSponsors(tt.args.ctx, tt.args.hackathon, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHackathonSponsors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHackathonSponsors() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetHackathonSponsors() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseRepository_GetHackathons(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx    context.Context
		filter *model.HackathonFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetHackathons(tt.args.ctx, tt.args.filter)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx context.Context
		obj *model.Sponsor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetHackathonsBySponsor(tt.args.ctx, tt.args.obj)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		queryable   database.Queryable
		hackathonID string
		userID      string
		status      model.ApplicationStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			if err := r.UpdateApplicantStatus(tt.args.ctx, tt.args.queryable, tt.args.hackathonID, tt.args.userID, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateApplicantStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateApplication(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx         context.Context
		hackathonID string
		userID      string
		input       model.HackathonApplicationInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.HackathonApplication
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.UpdateApplication(tt.args.ctx, tt.args.hackathonID, tt.args.userID, tt.args.input)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx   context.Context
		id    string
		input *model.HackathonUpdateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Hackathon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.UpdateHackathon(tt.args.ctx, tt.args.id, tt.args.input)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx       context.Context
		queryable database.Queryable
		id        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Term
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetTermById(tt.args.ctx, tt.args.queryable, tt.args.id)
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
	type fields struct {
		DatabasePool *pgxpool.Pool
		TermBiMap    *structure.BiMap
	}
	type args struct {
		ctx          context.Context
		queryable    database.Queryable
		termYear     int
		termSemester model.Semester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
				TermBiMap:    tt.fields.TermBiMap,
			}
			got, err := r.GetTermId(tt.args.ctx, tt.args.queryable, tt.args.termYear, tt.args.termSemester)
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
	tests := []struct {
		name string
		args args
		want *repository.DatabaseRepository
	}{
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
