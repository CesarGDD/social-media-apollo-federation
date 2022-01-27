package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/users-social-media/graph/generated"
	"cesargdd/users-social-media/pg"
	"context"
	"fmt"
)

func (r *commentResolver) User(ctx context.Context, obj *pg.Comment) (*pg.User, error) {
	user, err := db.GetUserByID(context.Background(), int32(obj.UserID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *commentLikeResolver) User(ctx context.Context, obj *pg.CommentLike) (*pg.User, error) {
	user, err := db.GetUserByID(context.Background(), int32(obj.UserID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *followerResolver) FollowerID(ctx context.Context, obj *pg.Follower) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *followerResolver) User(ctx context.Context, obj *pg.Follower) (*pg.User, error) {
	user, err := db.GetUserByID(ctx, int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *pg.NewUser) (*pg.User, error) {
	createUser, err := db.CreateUser(context.Background(), pg.CreateUserParams{
		Username: input.Username,
		Bio:      *input.Bio,
		Avatar:   *input.Avatar,
		Email:    *input.Email,
		Password: input.Password,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        createUser.ID,
		CreatedAt: createUser.CreatedAt,
		UpdatedAt: createUser.UpdatedAt,
		Username:  createUser.Username,
		Bio:       createUser.Bio,
		Avatar:    createUser.Avatar,
		Email:     createUser.Email,
	}, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *pg.EditUser) (*pg.User, error) {
	updateUser, err := db.UpdateUser(context.Background(), pg.UpdateUserParams{
		ID:     int32(input.ID),
		Bio:    *input.Bio,
		Avatar: *input.Avatar,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        updateUser.ID,
		CreatedAt: updateUser.CreatedAt,
		UpdatedAt: updateUser.UpdatedAt,
		Username:  updateUser.Username,
		Bio:       updateUser.Bio,
		Avatar:    updateUser.Avatar,
		Email:     updateUser.Email,
	}, err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *int) (*pg.User, error) {
	deleteUser, err := db.DeleteUser(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteUser, err
}

func (r *postResolver) User(ctx context.Context, obj *pg.Post) (*pg.User, error) {
	user, err := db.GetUserByID(context.Background(), int32(obj.UserID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *postLikeResolver) User(ctx context.Context, obj *pg.PostLike) (*pg.User, error) {
	user, err := db.GetUserByID(context.Background(), int32(obj.UserID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*pg.User, error) {
	user, err := db.GetUserByUsername(context.Background(), username)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *queryResolver) UserByID(ctx context.Context, id int) (*pg.User, error) {
	user, err := db.GetUserByID(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, err
}

func (r *queryResolver) Users(ctx context.Context) ([]pg.User, error) {
	users, err := db.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return users, err
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// CommentLike returns generated.CommentLikeResolver implementation.
func (r *Resolver) CommentLike() generated.CommentLikeResolver { return &commentLikeResolver{r} }

// Follower returns generated.FollowerResolver implementation.
func (r *Resolver) Follower() generated.FollowerResolver { return &followerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
type commentLikeResolver struct{ *Resolver }
type followerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type postLikeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *followerResolver) LeaderID(ctx context.Context, obj *pg.Follower) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
