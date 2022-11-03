package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_shared/pagination"
)

// Hackathon is the resolver for the hackathon field.
func (r *eventResolver) Hackathon(ctx context.Context, obj *model.Event) (*model.Hackathon, error) {
	return r.Repository.GetHackathonByEvent(ctx, obj)
}

// Applicants is the resolver for the applicants field.
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

// Attendees is the resolver for the attendees field.
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

// Sponsors is the resolver for the sponsors field.
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

// Events is the resolver for the events field.
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

// Status is the resolver for the status field.
func (r *hackathonResolver) Status(ctx context.Context, obj *model.Hackathon) (model.HackathonStatus, error) {
	now := time.Now().UTC()

	if obj.StartDate.After(now) {
		return model.HackathonStatusFuture, nil
	}
	if obj.EndDate.Before(now) {
		return model.HackathonStatusPast, nil
	}

	return model.HackathonStatusPresent, nil
}

// Pending is the resolver for the pending field.
func (r *hackathonResolver) Pending(ctx context.Context, obj *model.Hackathon, userID string) (bool, error) {
	return r.Repository.IsUserPending(ctx, obj, userID)
}

// Attending is the resolver for the attending field.
func (r *hackathonResolver) Attending(ctx context.Context, obj *model.Hackathon, userID string) (bool, error) {
	return r.Repository.IsUserAttending(ctx, obj, userID)
}

// CreateHackathon is the resolver for the createHackathon field.
func (r *mutationResolver) CreateHackathon(ctx context.Context, input model.HackathonCreateInput) (*model.Hackathon, error) {
	return r.Repository.CreateHackathon(ctx, &input)
}

// UpdateHackathon is the resolver for the updateHackathon field.
func (r *mutationResolver) UpdateHackathon(ctx context.Context, id string, input model.HackathonUpdateInput) (*model.Hackathon, error) {
	return r.Repository.UpdateHackathon(ctx, id, &input)
}

// DeleteHackathon is the resolver for the deleteHackathon field.
func (r *mutationResolver) DeleteHackathon(ctx context.Context, id string) (bool, error) {
	return r.Repository.DeleteHackathon(ctx, id)
}

// AcceptApplicant is the resolver for the acceptApplicant field.
func (r *mutationResolver) AcceptApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	return r.Repository.AcceptApplicant(ctx, hackathonID, userID)
}

// DenyApplicant is the resolver for the denyApplicant field.
func (r *mutationResolver) DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
	return r.Repository.DenyApplicant(ctx, hackathonID, userID)
}

// ApplyToHackathon is the resolver for the applyToHackathon field.
func (r *mutationResolver) ApplyToHackathon(ctx context.Context, hackathonID string, input model.HackathonApplicationInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// CurrentHackathon is the resolver for the currentHackathon field.
func (r *queryResolver) CurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	return r.Repository.GetCurrentHackathon(ctx)
}

// Hackathons is the resolver for the hackathons field.
func (r *queryResolver) Hackathons(ctx context.Context, filter model.HackathonFilter) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathons(ctx, &filter)
}

// GetHackathon is the resolver for the getHackathon field.
func (r *queryResolver) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	return r.Repository.GetHackathon(ctx, id)
}

// Hackathons is the resolver for the hackathons field.
func (r *sponsorResolver) Hackathons(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsBySponsor(ctx, obj)
}

// AttendedHackathons is the resolver for the attendedHackathons field.
func (r *userResolver) AttendedHackathons(ctx context.Context, obj *model.User) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsByUser(ctx, obj, true)
}

// AppliedHackathons is the resolver for the appliedHackathons field.
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
