package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/posts-social-media/graph/generated"
	"cesargdd/posts-social-media/pg"
	"context"
	"fmt"
)

func (r *hashtagPostResolver) Posts(ctx context.Context, obj *pg.HashtagPost) ([]pg.Post, error) {
	posts, err := db.ListPostsById(context.Background(), int32(obj.PostID))
	if err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *pg.NewPost) (*pg.Post, error) {
	createPost, err := db.CreatePost(context.Background(), pg.CreatePostParams{
		Url:     input.URL,
		Caption: *input.Caption,
		UserID:  int32(input.UserID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        createPost.ID,
		CreatedAt: createPost.CreatedAt,
		UpdatedAt: createPost.UpdatedAt,
		URL:       createPost.URL,
		Caption:   createPost.Caption,
		UserID:    createPost.UserID,
	}, err
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input *pg.EditPost) (*pg.Post, error) {
	updatePost, err := db.UpdatePost(context.Background(), pg.UpdatePostParams{
		ID:      int32(input.ID),
		Url:     *input.URL,
		Caption: *input.Caption,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        updatePost.ID,
		CreatedAt: updatePost.CreatedAt,
		UpdatedAt: updatePost.UpdatedAt,
		URL:       updatePost.URL,
		Caption:   updatePost.Caption,
		UserID:    updatePost.UserID,
	}, err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *int) (*pg.Post, error) {
	deletePost, err := db.DeletePost(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deletePost, err
}

func (r *queryResolver) Post(ctx context.Context, id int) (*pg.Post, error) {
	post, err := db.GetPost(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		URL:       post.URL,
		Caption:   post.Caption,
		UserID:    post.UserID,
	}, err
}

func (r *queryResolver) Posts(ctx context.Context) ([]pg.Post, error) {
	posts, err := db.ListPosts(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

func (r *userResolver) Posts(ctx context.Context, obj *pg.User) ([]pg.Post, error) {
	posts, err := db.ListPostsUser(context.Background(), int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

// HashtagPost returns generated.HashtagPostResolver implementation.
func (r *Resolver) HashtagPost() generated.HashtagPostResolver { return &hashtagPostResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type hashtagPostResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

var conn = pg.Connect()
var db = pg.New(conn)
