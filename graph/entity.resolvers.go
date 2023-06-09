package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
)

func (r *entityResolver) FindEventByID(ctx context.Context, id string) (*model.Event, error) {
	return &model.Event{ID: id}, nil
}

func (r *entityResolver) FindHackathonByID(ctx context.Context, id string) (*model.Hackathon, error) {
	return r.Repository.GetHackathon(ctx, id)
}

func (r *entityResolver) FindHackathonByTermYearAndTermSemester(ctx context.Context, termYear int, termSemester model.Semester) (*model.Hackathon, error) {
	return r.Repository.GetHackathonByTermYearAndTermSemester(ctx, termYear, termSemester)
}

func (r *entityResolver) FindHackathonApplicationByID(ctx context.Context, id string) (*model.HackathonApplication, error) {
	split := strings.Split(id, "-")
	return r.Repository.GetApplication(ctx, split[0], split[1])
}

func (r *entityResolver) FindSponsorByID(ctx context.Context, id string) (*model.Sponsor, error) {
	return &model.Sponsor{ID: id}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{ID: id}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
