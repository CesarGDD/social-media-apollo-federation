package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/comments-social-media/graph/generated"
	"cesargdd/comments-social-media/pg"
	"context"
)

func (r *entityResolver) FindCommentByPostID(ctx context.Context, postID int) (*pg.Comment, error) {
	return &pg.Comment{
		PostID: postID,
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
