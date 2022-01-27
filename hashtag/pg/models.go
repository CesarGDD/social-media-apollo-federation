// Code generated by sqlc. DO NOT EDIT.

package pg

import (
	"time"
)

type EditHashtag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Hashtag struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
}

func (Hashtag) IsEntity() {}

type HashtagPost struct {
	ID        int `json:"id"`
	HashtagID int `json:"hashtag_id"`
	PostID    int `json:"post_id"`
}

func (HashtagPost) IsEntity() {}

type NewHashtag struct {
	Title string `json:"title"`
}

type NewHashtagPost struct {
	HashtagID int `json:"hashtag_id"`
	PostID    int `json:"post_id"`
}

type Post struct {
	ID       int              `json:"id"`
	Hashtags []HashtagPost `json:"hashtags"`
}

func (Post) IsEntity() {}
