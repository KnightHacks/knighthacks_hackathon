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
	GetHackathonByEvent(ctx context.Context, obj *model.Event) (*model.Hackathon, error)

	DeleteHackathon(ctx context.Context, id string) (bool, error)

	GetCurrentHackathon(ctx context.Context) (*model.Hackathon, error)

	AcceptApplicant(ctx context.Context, hackathonID string, userID string) (bool, error)
	DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error)
	// TODO: DELETE IsUserAttending & IsUserPending
	IsUserAttending(ctx context.Context, hackathon *model.Hackathon, userID string) (bool, error)
	IsUserPending(ctx context.Context, hackathon *model.Hackathon, userID string) (bool, error)

	// Array returns

	GetHackathons(ctx context.Context, filter *model.HackathonFilter) ([]*model.Hackathon, error)
	// TODO: DELETE GetHackathonsByUser
	GetHackathonsByUser(ctx context.Context, obj *model.User, attended bool) ([]*model.Hackathon, error)
	GetHackathonsBySponsor(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error)

	// TODO: DELETE GetHackathonApplicants & GetHackathonAttendees

	GetHackathonApplicants(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.User, int, error)
	GetHackathonAttendees(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.User, int, error)
	GetHackathonSponsors(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Sponsor, int, error)
	GetHackathonEvents(ctx context.Context, hackathon *model.Hackathon, first int, after string) ([]*model.Event, int, error)

	GetApplicationsByUser(ctx context.Context, obj *model.User) ([]*model.HackathonApplication, error)
	GetApplication(ctx context.Context, hackathonID string, userID string) (*model.HackathonApplication, error)
	ApplyToHackathon(ctx context.Context, hackathonID string, input model.HackathonApplicationInput) (bool, error)
	UpdateApplication(ctx context.Context, hackathonID string, userID string, input model.HackathonApplicationInput) (*model.HackathonApplication, error)
	GetApplicationsByHackathon(ctx context.Context, obj *model.Hackathon, first int, after *string, status model.ApplicationStatus) ([]*model.HackathonApplication, int, error)
}
