package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/followers-social-media/graph/generated"
	"cesargdd/followers-social-media/pg"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateFollower(ctx context.Context, input *pg.NewFollower) (*pg.Follower, error) {
	createFollower, err := db.CreateFollower(ctx, pg.CreateFollowerParams{
		LeaderID:   int32(input.LeaderID),
		FollowerID: int32(input.FollowerID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Follower{
		ID:         createFollower.ID,
		CreatedAt:  createFollower.CreatedAt,
		LeaderID:   createFollower.LeaderID,
		FollowerID: createFollower.FollowerID,
	}, nil
}

func (r *mutationResolver) DeleteFollower(ctx context.Context, input *int) (*pg.Follower, error) {
	deleteFollower, err := db.DeleteFollower(ctx, int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteFollower, nil
}

func (r *queryResolver) Follower(ctx context.Context, id int) (*pg.Follower, error) {
	follower, err := db.GetFollower(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Follower{
		ID:         follower.ID,
		CreatedAt:  follower.CreatedAt,
		LeaderID:   follower.LeaderID,
		FollowerID: follower.FollowerID,
	}, nil
}

func (r *queryResolver) Followers(ctx context.Context) ([]pg.Follower, error) {
	followers, err := db.ListFollowers(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return followers, nil
}

func (r *userResolver) Followers(ctx context.Context, obj *pg.User) ([]pg.Follower, error) {
	user, err := db.ListFollowersByLeaderId(ctx, int32(obj.ID))
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

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
