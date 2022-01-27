package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/followers-social-media/graph/generated"
	"cesargdd/followers-social-media/pg"
	"context"
)

func (r *entityResolver) FindFollowerByID(ctx context.Context, id int) (*pg.Follower, error) {
	return &pg.Follower{
		ID: id,
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*pg.User, error) {
	return &pg.User{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
