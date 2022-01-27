package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/users-social-media/graph/generated"
	"cesargdd/users-social-media/pg"
	"context"
	"fmt"
)

func (r *entityResolver) FindCommentByUserID(ctx context.Context, userID int) (*pg.Comment, error) {
	return &pg.Comment{
		UserID: userID,
	}, nil
}

func (r *entityResolver) FindCommentLikeByUserID(ctx context.Context, userID int) (*pg.CommentLike, error) {
	return &pg.CommentLike{
		UserID: userID,
	}, nil
}

func (r *entityResolver) FindFollowerByFollowerID(ctx context.Context, followerID int) (*pg.Follower, error) {
	return &pg.Follower{
		ID: followerID,
	}, nil
}

func (r *entityResolver) FindPostByUserID(ctx context.Context, userID int) (*pg.Post, error) {
	return &pg.Post{
		UserID: userID,
	}, nil
}

func (r *entityResolver) FindPostLikeByUserID(ctx context.Context, userID int) (*pg.PostLike, error) {
	return &pg.PostLike{
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *entityResolver) FindFollowerByLeaderID(ctx context.Context, leaderID int) (*pg.Follower, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *entityResolver) FindFollowerByID(ctx context.Context, id int) (*pg.Follower, error) {
	return &pg.Follower{
		ID: id,
	}, nil
}

var conn = pg.Connect()
var db = pg.New(conn)
