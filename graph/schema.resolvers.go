package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

func (r *hackathonResolver) Status(ctx context.Context, obj *model.Hackathon) (model.HackathonStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateHackathon(ctx context.Context, input model.HackathonCreateInput) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateHackathon(ctx context.Context, input model.HackathonUpdateInput) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hackathons(ctx context.Context, filter model.HackathonFilter) ([]*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

// Hackathon returns generated.HackathonResolver implementation.
func (r *Resolver) Hackathon() generated.HackathonResolver { return &hackathonResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type hackathonResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
