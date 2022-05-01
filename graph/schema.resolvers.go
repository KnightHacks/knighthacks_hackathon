package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/LockedThread/knighthacks_hackathon/graph/generated"
	"github.com/LockedThread/knighthacks_hackathon/graph/model"
)

func (r *queryResolver) CurrentHackathon(ctx context.Context) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hackathons(ctx context.Context, input model.HackathonFilterInput) ([]*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHackathon(ctx context.Context, id string) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
