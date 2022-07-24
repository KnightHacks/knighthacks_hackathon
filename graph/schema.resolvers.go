package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

func (r *eventResolver) Hackathon(ctx context.Context, obj *model.Event) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Applicants(ctx context.Context, obj *model.Hackathon) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Attendees(ctx context.Context, obj *model.Hackathon) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Sponsors(ctx context.Context, obj *model.Hackathon) ([]*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Events(ctx context.Context, obj *model.Hackathon) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Status(ctx context.Context, obj *model.Hackathon) (model.HackathonStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *hackathonResolver) Attending(ctx context.Context, obj *model.Hackathon, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DenyApplicant(ctx context.Context, hackathonID string, userID string) (bool, error) {
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
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) AttendedHackathons(ctx context.Context, obj *model.User) ([]*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) AppliedHackathons(ctx context.Context, obj *model.User) ([]*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
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
