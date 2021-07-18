package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dacharat/go-gql/graph/generated"
	"github.com/dacharat/go-gql/graph/model"
	"github.com/dacharat/go-gql/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	link.Address = input.Address
	link.Title = input.Title
	var user model.User
	user.Name = "test"
	link.User = &user
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user model.User
	user.Name = "test"
	return user.ID, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return jwt.GenerateToken(input.Username)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	// panic(fmt.Errorf("not implemented"))
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	links = append(links, &model.Link{Title: "our dummy link", Address: "https://address.org", User: &model.User{Name: "admin"}})
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
