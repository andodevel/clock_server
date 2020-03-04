package graphql

import (
	"context"

	"github.com/jinzhu/gorm"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver ...
type Resolver struct {
	db *gorm.DB
}

// NewGormConfig ...
func NewGormConfig(db *gorm.DB) Config {
	return Config{Resolvers: &Resolver{
		db: db,
	}}
}

// Mutation ...
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query ...
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Register(ctx context.Context, username string, password string) (*User, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*User, error) {
	panic("not implemented")
}
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) App(ctx context.Context) (*App, error) {
	var app = new(App)
	app.Name = "Codelink SSO Proof Of Concept"
	app.Version = "0.0.1"
	return app, nil
}
func (r *queryResolver) User(ctx context.Context, id *int) (*User, error) {
	var user User
	r.db.First(&user, *id)
	return &user, nil
}
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	var users []*User
	r.db.Find(&users)
	return users, nil
}
