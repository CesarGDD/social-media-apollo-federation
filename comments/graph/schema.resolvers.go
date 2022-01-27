package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/comments-social-media/graph/generated"
	"cesargdd/comments-social-media/pg"
	"context"
	"fmt"
)

func (r *commentLikeResolver) CommentID(ctx context.Context, obj *pg.CommentLike) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *pg.NewComment) (*pg.Comment, error) {
	createComment, err := db.CreateComment(context.Background(), pg.CreateCommentParams{
		Contents: input.Contents,
		UserID:   int32(input.UserID),
		PostID:   int32(input.PostID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        createComment.ID,
		CreatedAt: createComment.CreatedAt,
		UpdatedAt: createComment.UpdatedAt,
		Contents:  createComment.Contents,
		UserID:    createComment.UserID,
		PostID:    createComment.PostID,
	}, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, input *pg.EditComment) (*pg.Comment, error) {
	updateComment, err := db.UpdateComment(context.Background(), pg.UpdateCommentParams{
		ID:       int32(input.ID),
		Contents: input.Contents,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        updateComment.ID,
		CreatedAt: updateComment.CreatedAt,
		UpdatedAt: updateComment.UpdatedAt,
		Contents:  updateComment.Contents,
		UserID:    updateComment.UserID,
		PostID:    updateComment.PostID,
	}, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, input *int) (*pg.Comment, error) {
	deleteComment, err := db.DeleteComment(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteComment, nil
}

func (r *postResolver) Comments(ctx context.Context, obj *pg.Post) ([]pg.Comment, error) {
	commentsPost, err := db.ListCommentsPost(context.Background(), int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return commentsPost, nil
}

func (r *queryResolver) Comment(ctx context.Context, id int) (*pg.Comment, error) {
	comment, err := db.GetComment(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		UserID:    comment.UserID,
		PostID:    comment.PostID,
	}, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]pg.Comment, error) {
	comments, err := db.ListComments(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return comments, nil
}

// CommentLike returns generated.CommentLikeResolver implementation.
func (r *Resolver) CommentLike() generated.CommentLikeResolver { return &commentLikeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type commentLikeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var conn = pg.Connect()
var db = pg.New(conn)
