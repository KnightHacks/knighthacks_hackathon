package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_shared/pagination"
)

func (r *eventResolver) Hackathon(ctx context.Context, obj *model.Event) (*model.Hackathon, error) {
	return r.Repository.GetHackathonByEvent(ctx, obj)
}

func (r *hackathonResolver) Applicants(ctx context.Context, obj *model.Hackathon, first int, after *string) (*model.UsersConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	applicants, total, err := r.Repository.GetHackathonApplicants(ctx, obj, first, a)
	if err != nil {
		return nil, err
	}
	connection := model.UsersConnection{Users: applicants,
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(applicants[0].ID, applicants[len(applicants)-1].ID),
	}
	return &connection, nil
}

func (r *hackathonResolver) Attendees(ctx context.Context, obj *model.Hackathon, first int, after *string) (*model.UsersConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	attendees, total, err := r.Repository.GetHackathonAttendees(ctx, obj, first, a)
	if err != nil {
		return nil, err
	}
	connection := model.UsersConnection{Users: attendees,
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(attendees[0].ID, attendees[len(attendees)-1].ID),
	}
	return &connection, err
}

func (r *hackathonResolver) Sponsors(ctx context.Context, obj *model.Hackathon, first int, after *string) (*model.SponsorsConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	sponsors, total, err := r.Repository.GetHackathonSponsors(ctx, obj, first, a)
	if err != nil {
		return nil, err
	}
	connection := model.SponsorsConnection{Sponsors: sponsors,
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(sponsors[0].ID, sponsors[len(sponsors)-1].ID),
	}
	return &connection, err
}

func (r *hackathonResolver) Events(ctx context.Context, obj *model.Hackathon, first int, after *string) (*model.EventsConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	events, total, err := r.Repository.GetHackathonEvents(ctx, obj, first, a)
	if err != nil {
		return nil, err
	}
	connection := model.EventsConnection{Events: events,
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(events[0].ID, events[len(events)-1].ID),
	}
	return &connection, err
}

func (r *hackathonResolver) Status(ctx context.Context, obj *model.Hackathon) (model.HackathonStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Pending(ctx context.Context, obj *model.Hackathon, userID string) (bool, error) {
	return r.Repository.IsUserPending(ctx, obj, userID)
}

func (r *hackathonResolver) Attending(ctx context.Context, obj *model.Hackathon, userID string) (bool, error) {
	return r.Repository.IsUserAttending(ctx, obj, userID)
}

func (r *mutationResolver) CreateHackathon(ctx context.Context, input model.HackathonCreateInput) (*model.Hackathon, error) {
	return r.Repository.CreateHackathon(ctx, &input)
}

func (r *mutationResolver) UpdateHackathon(ctx context.Context, id string, input model.HackathonUpdateInput) (*model.Hackathon, error) {
	return r.Repository.UpdateHackathon(ctx, id, &input)
}

func (r *mutationResolver) DeleteHackathon(ctx context.Context, id string) (bool, error) {
	return r.Repository.DeleteHackathon(ctx, id)
}

func (r *mutationResolver) AcceptApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	return r.Repository.AcceptApplicant(ctx, hackathonID, userID)
}

func (r *mutationResolver) DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	return r.Repository.DenyApplicant(ctx, hackathonID, userID)
}

func (r *mutationResolver) ApplyToHackathon(ctx context.Context, hackathonID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	return r.Repository.GetCurrentHackathon(ctx)
}

func (r *queryResolver) Hackathons(ctx context.Context, filter model.HackathonFilter) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathons(ctx, &filter)
}

func (r *queryResolver) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	return r.Repository.GetHackathon(ctx, id)
}

func (r *sponsorResolver) Hackathons(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsBySponsor(ctx, obj)
}

func (r *userResolver) AttendedHackathons(ctx context.Context, obj *model.User) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsByUser(ctx, obj, true)
}

func (r *userResolver) AppliedHackathons(ctx context.Context, obj *model.User) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsByUser(ctx, obj, false)
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Hackathon returns generated.HackathonResolver implementation.
func (r *Resolver) Hackathon() generated.HackathonResolver { return &hackathonResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Sponsor returns generated.SponsorResolver implementation.
func (r *Resolver) Sponsor() generated.SponsorResolver { return &sponsorResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type eventResolver struct{ *Resolver }
type hackathonResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type sponsorResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
