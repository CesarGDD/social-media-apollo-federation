package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/likes-social-media/graph/generated"
	"cesargdd/likes-social-media/pg"
	"context"
	"fmt"
)

func (r *entityResolver) FindCommentByID(ctx context.Context, id int) (*pg.Comment, error) {
	return &pg.Comment{
		ID: id,
	}, nil
}

func (r *entityResolver) FindCommentLikeByUserID(ctx context.Context, userID int) (*pg.CommentLike, error) {
	return &pg.CommentLike{
		UserID: userID,
	}, nil
}

func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*pg.Post, error) {
	return &pg.Post{
		ID: id,
	}, nil
}

func (r *entityResolver) FindPostLikeByUserID(ctx context.Context, userID int) (*pg.PostLike, error) {
	return &pg.PostLike{
		UserID: userID,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *entityResolver) FindCommentLikeByID(ctx context.Context, id int) (*pg.CommentLike, error) {
	comment, err := db.GetCommentLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.CommentLike{
		ID:        comment.CommentID,
		CreatedAt: comment.CreatedAt,
		UserID:    comment.UserID,
		CommentID: comment.CommentID,
	}, nil
}
func (r *entityResolver) FindPostLikeByID(ctx context.Context, id int) (*pg.PostLike, error) {
	post, err := db.GetPostLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.PostLike{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UserID:    post.UserID,
		PostID:    post.PostID,
	}, nil
}
