package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/graph/model"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/models"
	"github.com/KnightHacks/knighthacks_shared/pagination"
)

func (r *eventResolver) Hackathon(ctx context.Context, obj *model.Event) (*model.Hackathon, error) {
	return r.Repository.GetHackathonByEvent(ctx, obj)
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
	now := time.Now().UTC()

	if obj.StartDate.After(now) {
		return model.HackathonStatusFuture, nil
	}
	if obj.EndDate.Before(now) {
		return model.HackathonStatusPast, nil
	}

	return model.HackathonStatusPresent, nil
}

func (r *hackathonResolver) Applications(ctx context.Context, obj *model.Hackathon, first int, after *string, status model.ApplicationStatus) (*model.HackathonApplicationConnection, error) {
	hackathons, total, err := r.Repository.GetApplicationsByHackathon(ctx, obj, first, after, status)

	connection := model.HackathonApplicationConnection{Applications: hackathons,
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(hackathons[0].ID, hackathons[len(hackathons)-1].ID),
	}
	return &connection, err
}

func (r *hackathonApplicationResolver) Hackathon(ctx context.Context, obj *model.HackathonApplication) (*model.Hackathon, error) {
	return r.Repository.GetHackathon(ctx, obj.ID)
}

func (r *hackathonApplicationResolver) ResumeBase64(ctx context.Context, obj *model.HackathonApplication) (*string, error) {
	if obj.ResumeBase64 != nil {
		return obj.ResumeBase64, nil
	}
	resume, err := r.AzureBlobClient.DownloadResume(ctx, obj.Hackathon.ID, obj.User.ID)
	if err != nil {
		return nil, err
	}

	resumeBase64Encoding := base64.StdEncoding.EncodeToString(resume)

	return &resumeBase64Encoding, nil
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

func (r *mutationResolver) UpdateApplication(ctx context.Context, hackathonID string, userID string, input model.HackathonApplicationInput) (*model.HackathonApplication, error) {
	claims, ok := ctx.Value("AuthorizationUserClaims").(*auth.UserClaims)
	if !ok {
		return nil, errors.New("unable to retrieve user claims, most likely forgot to set @hasRole directive")
	}
	if claims.Role != models.RoleAdmin && claims.Id != userID {
		return nil, errors.New("unauthorized to update hackathon application that is not you")
	}

	_, err := r.Repository.GetHackathon(ctx, hackathonID)
	if err != nil {
		return nil, err
	}

	var bytes []byte
	var application *model.HackathonApplication
	if input.Resume != nil {
		if r.AzureBlobClient != nil {
			bytes, err = io.ReadAll(input.Resume.File)
			if err != nil {
				return nil, err
			}
			err = r.AzureBlobClient.UploadResume(ctx, hackathonID, userID, bytes)
			if err != nil {
				return nil, err
			}
		}
		defer func() {
			base64EncodedFile := base64.StdEncoding.EncodeToString(bytes)
			application.ResumeBase64 = &base64EncodedFile
		}()
	}

	application, err = r.Repository.UpdateApplication(ctx, hackathonID, userID, input)
	if err != nil {
		return nil, err
	}
	application.UserID = userID
	application.HackathonID = hackathonID
	return application, err
}

func (r *mutationResolver) ApplyToHackathon(ctx context.Context, hackathonID string, input model.HackathonApplicationInput) (bool, error) {
	claims, ok := ctx.Value("AuthorizationUserClaims").(*auth.UserClaims)
	if !ok {
		return false, errors.New("unable to retrieve user claims, most likely forgot to set @hasRole directive")
	}

	return r.Repository.ApplyToHackathon(ctx, hackathonID, claims.UserID, input)
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

func (r *queryResolver) GetApplication(ctx context.Context, hackathonID string, userID string) (*model.HackathonApplication, error) {
	claims, ok := ctx.Value("AuthorizationUserClaims").(*auth.UserClaims)
	if !ok {
		return nil, errors.New("unable to retrieve user claims, most likely forgot to set @hasRole directive")
	}
	if claims.Role != models.RoleAdmin && claims.Id != userID {
		return nil, errors.New("unauthorized to update hackathon application that is not you")
	}

	return r.Entity().FindHackathonApplicationByID(ctx, fmt.Sprintf("%s-%s", hackathonID, userID))
}

func (r *sponsorResolver) Hackathons(ctx context.Context, obj *model.Sponsor) ([]*model.Hackathon, error) {
	return r.Repository.GetHackathonsBySponsor(ctx, obj)
}

func (r *userResolver) Applications(ctx context.Context, obj *model.User) ([]*model.HackathonApplication, error) {
	return r.Repository.GetApplicationsByUser(ctx, obj)
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Hackathon returns generated.HackathonResolver implementation.
func (r *Resolver) Hackathon() generated.HackathonResolver { return &hackathonResolver{r} }

// HackathonApplication returns generated.HackathonApplicationResolver implementation.
func (r *Resolver) HackathonApplication() generated.HackathonApplicationResolver {
	return &hackathonApplicationResolver{r}
}

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
type hackathonApplicationResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type sponsorResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *hackathonApplicationResolver) User(ctx context.Context, obj *model.HackathonApplication) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
