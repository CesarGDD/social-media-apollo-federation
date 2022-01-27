package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/likes-social-media/graph/generated"
	"cesargdd/likes-social-media/pg"
	"context"
	"fmt"
)

func (r *commentResolver) CommentLikes(ctx context.Context, obj *pg.Comment) ([]pg.CommentLike, error) {
	likes, err := db.ListLikesFromComment(context.Background(), int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return likes, err
}

func (r *mutationResolver) CreatePostLike(ctx context.Context, input *pg.NewPostLike) (*pg.PostLike, error) {
	createPostLike, err := db.CreatePostLike(context.Background(), pg.CreatePostLikeParams{
		UserID: int32(input.UserID),
		PostID: int32(input.PostID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.PostLike{
		ID:        createPostLike.ID,
		CreatedAt: createPostLike.CreatedAt,
		UserID:    createPostLike.UserID,
		PostID:    createPostLike.PostID,
	}, nil
}

func (r *mutationResolver) DeletePostLike(ctx context.Context, input *int) (*pg.PostLike, error) {
	deletePostLike, err := db.DeletePostLike(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deletePostLike, nil
}

func (r *mutationResolver) CreateCommentLike(ctx context.Context, input *pg.NewCommentLike) (*pg.CommentLike, error) {
	createCommentLike, err := db.CreateCommentLike(context.Background(), pg.CreateCommentLikeParams{
		UserID:    int32(input.UserID),
		CommentID: int32(input.CommentID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.CommentLike{
		ID:        createCommentLike.ID,
		CreatedAt: createCommentLike.CreatedAt,
		UserID:    createCommentLike.UserID,
		CommentID: createCommentLike.CommentID,
	}, nil
}

func (r *mutationResolver) DeleteCommentLike(ctx context.Context, input *int) (*pg.CommentLike, error) {
	deleteCommentLike, err := db.DeleteCommentLike(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteCommentLike, nil
}

func (r *postResolver) PostLikes(ctx context.Context, obj *pg.Post) ([]pg.PostLike, error) {
	likes, err := db.ListLikesFromPost(context.Background(), int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return likes, nil
}

func (r *queryResolver) PostLike(ctx context.Context, id int) (*pg.PostLike, error) {
	postLike, err := db.GetPostLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.PostLike{
		ID:        postLike.ID,
		CreatedAt: postLike.CreatedAt,
		UserID:    postLike.UserID,
		PostID:    postLike.PostID,
	}, nil
}

func (r *queryResolver) PostLikes(ctx context.Context) ([]pg.PostLike, error) {
	postsLikes, err := db.ListPostLikes(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return postsLikes, nil
}

func (r *queryResolver) CommentLike(ctx context.Context, id int) (*pg.CommentLike, error) {
	commentLike, err := db.GetCommentLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.CommentLike{
		ID:        commentLike.ID,
		CreatedAt: commentLike.CreatedAt,
		UserID:    commentLike.UserID,
		CommentID: commentLike.CommentID,
	}, nil
}

func (r *queryResolver) CommentLikes(ctx context.Context) ([]pg.CommentLike, error) {
	commentLikes, err := db.ListCommentLikes(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return commentLikes, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
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
