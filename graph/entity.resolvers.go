package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/LockedThread/knighthacks_hackathon/graph/generated"
	"github.com/LockedThread/knighthacks_hackathon/graph/model"
)

func (r *entityResolver) FindEventByID(ctx context.Context, id string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindHackathonByID(ctx context.Context, id string) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindSponsorByID(ctx context.Context, id string) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
