package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

type Repository interface {
	CreateHackathon(ctx context.Context, input *model.HackathonCreateInput) (*model.Hackathon, error)
	UpdateHackathon(ctx context.Context, id int, input *model.HackathonUpdateInput) (*model.Hackathon, error)
	GetHackathon(ctx context.Context, id int) (*model.Hackathon, error)
	GetHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error)
	GetHackathonByEvent(ctx context.Context, obj *model.Event) (*model.Hackathon, error)

	DeleteHackathon(ctx context.Context, id int) (bool, error)

	GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error)

	AcceptApplicant(ctx context.Context, hackathonid int, userid int) (bool, error)
	DenyApplicant(ctx context.Context, hackathonid int, userid int) (bool, error)
	// Array returns

	GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error)
	GetHackathonsBySponsor(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error)

	GetHackathonSponsors(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Sponsor, int, error)
	GetHackathonEvents(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Event, int, error)

	GetApplicationsByUser(ctx context.Context, obj *model.User) ([]*model.HackathonApplication, error)
	GetApplication(ctx context.Context, hackathonid int, userid int) (*model.HackathonApplication, error)
	ApplyToHackathon(ctx context.Context, hackathonid int, userid int, input model.HackathonApplicationInput) (bool, error)
	UpdateApplication(ctx context.Context, hackathonid int, userid int, input model.HackathonApplicationInput) (*model.HackathonApplication, error)
	GetApplicationsByHackathon(ctx context.Context, obj *model.Hackathon, first int, after *string, status model.ApplicationStatus) ([]*model.HackathonApplication, int, error)
}
