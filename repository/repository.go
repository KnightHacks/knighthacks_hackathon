package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

type Repository interface {
	CreateHackathon(ctx context.Context, input model.HackathonCreateInput) (*model.Hackathon, error)
	UpdateHackathon(ctx context.Context, id string, input model.HackathonUpdateInput) (*model.Hackathon, error)
	GetHackathon(ctx context.Context, id string) (*model.Hackathon, error)
	GetHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error)

	DeleteHackathon(ctx context.Context, id string) (*model.Hackathon, error)

	GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error)

	// Array returns

	GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error)

	// Shared References

	GetEventById()
}
