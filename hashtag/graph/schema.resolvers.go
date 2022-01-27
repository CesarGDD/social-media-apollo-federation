package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/hashtag-social-media/graph/generated"
	"cesargdd/hashtag-social-media/pg"
	"context"
	"fmt"
)

var conn = pg.Connect()
var db = pg.New(conn)

// var ctx = context.Background()

func (r *hashtagPostResolver) Hashtag(ctx context.Context, obj *pg.HashtagPost) (*pg.Hashtag, error) {
	hashtag, err := db.GetHashtagById(ctx, int32(obj.PostID))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        hashtag.ID,
		Title:     hashtag.Title,
		CreatedAt: hashtag.CreatedAt,
	}, nil
}

func (r *mutationResolver) CreateHashtag(ctx context.Context, input *pg.NewHashtag) (*pg.Hashtag, error) {
	createHashtag, err := db.CreateHashtag(ctx, input.Title)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        createHashtag.ID,
		CreatedAt: createHashtag.CreatedAt,
		Title:     createHashtag.Title,
	}, nil
}

func (r *mutationResolver) UpdateHashtag(ctx context.Context, input *pg.EditHashtag) (*pg.Hashtag, error) {
	updateHashtag, err := db.UpdateHashtag(ctx, pg.UpdateHashtagParams{
		ID:    int32(input.ID),
		Title: input.Title,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        updateHashtag.ID,
		CreatedAt: updateHashtag.CreatedAt,
		Title:     updateHashtag.Title,
	}, nil
}

func (r *mutationResolver) DeleteHashtag(ctx context.Context, input *int) (*pg.Hashtag, error) {
	deleteHastagh, err := db.DeleteHashtag(ctx, int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteHastagh, nil
}

func (r *mutationResolver) CreateHashtagPost(ctx context.Context, input *pg.NewHashtagPost) (*pg.HashtagPost, error) {
	createHasPost, err := db.CreateHashtagPost(ctx, pg.CreateHashtagPostParams{
		PostID:    int32(input.PostID),
		HashtagID: int32(input.HashtagID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.HashtagPost{
		ID:        createHasPost.ID,
		HashtagID: createHasPost.HashtagID,
		PostID:    createHasPost.PostID,
	}, nil
}

func (r *mutationResolver) DeleteHashtagPost(ctx context.Context, input *int) (*pg.HashtagPost, error) {
	deleteHasPost, err := db.DeleteHashtagPost(ctx, int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteHasPost, nil
}

func (r *queryResolver) HashtagByTitle(ctx context.Context, title string) (*pg.Hashtag, error) {
	hashtag, err := db.GetHashtagByTitle(ctx, title)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        hashtag.ID,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *queryResolver) HashtagsByID(ctx context.Context, id int) (*pg.Hashtag, error) {
	hashtag, err := db.GetHashtagById(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        hashtag.ID,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *queryResolver) Hashtags(ctx context.Context) ([]pg.Hashtag, error) {
	hashtags, err := db.ListHashtags(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return hashtags, nil
}

func (r *queryResolver) HashtagPost(ctx context.Context, id int) (*pg.HashtagPost, error) {
	hasPost, err := db.GetHashtagPostById(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.HashtagPost{
		ID:        hasPost.ID,
		HashtagID: hasPost.HashtagID,
		PostID:    hasPost.PostID,
	}, nil
}

func (r *queryResolver) HashtagsPost(ctx context.Context) ([]pg.HashtagPost, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return hasPosts, nil
}

// HashtagPost returns generated.HashtagPostResolver implementation.
func (r *Resolver) HashtagPost() generated.HashtagPostResolver { return &hashtagPostResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type hashtagPostResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
