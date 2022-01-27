package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/posts-social-media/graph/generated"
	"cesargdd/posts-social-media/pg"
	"context"
)

func (r *entityResolver) FindHashtagPostByPostID(ctx context.Context, postID int) (*pg.HashtagPost, error) {
	return &pg.HashtagPost{
		PostID: postID,
	}, nil
}

func (r *entityResolver) FindPostByUserID(ctx context.Context, userID int) (*pg.Post, error) {
	return &pg.Post{
		UserID: userID,
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
