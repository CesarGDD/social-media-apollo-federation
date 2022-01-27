package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/hashtag-social-media/graph/generated"
	"cesargdd/hashtag-social-media/pg"
	"context"
)

func (r *entityResolver) FindHashtagByID(ctx context.Context, id int) (*pg.Hashtag, error) {
	return &pg.Hashtag{
		ID: id,
	}, nil
}

func (r *entityResolver) FindHashtagPostByID(ctx context.Context, id int) (*pg.HashtagPost, error) {
	return &pg.HashtagPost{
		ID: id,
	}, nil
}

func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*pg.Post, error) {
	return &pg.Post{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
