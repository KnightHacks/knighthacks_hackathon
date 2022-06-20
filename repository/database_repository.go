package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

//DatabaseRepository
//Implements the Repository interface's functions
type DatabaseRepository struct {
	DatabasePool *pgxpool.Pool
}

func NewDatabaseRepository(databasePool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePool: databasePool,
	}
}

func (r *DatabaseRepository) CreateHackathon(ctx context.Context, input model.HackathonCreateInput) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) UpdateHackathon(ctx context.Context, id string, input model.HackathonUpdateInput) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) DeleteHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetEventById() {
	//TODO implement me
	panic("implement me")
}
