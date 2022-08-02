package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

type Repository interface {
	CreateHackathon(ctx context.Context, input *model.HackathonCreateInput) (*model.Hackathon, error)
	UpdateHackathon(ctx context.Context, id string, input *model.HackathonUpdateInput) (*model.Hackathon, error)
	GetHackathon(ctx context.Context, id string) (*model.Hackathon, error)
	GetHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error)

	DeleteHackathon(ctx context.Context, id string) (bool, error)

	GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error)

	AcceptApplicant(ctx context.Context, hackathonID string, userID string) (bool, error)
	DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error)
	// Array returns

	GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error)

	GetHackathonsBySponsor(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error)
}
