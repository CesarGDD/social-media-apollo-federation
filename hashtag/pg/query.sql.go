// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

const createHashtag = `-- name: CreateHashtag :one
INSERT INTO hashtags (title)
VALUES ($1)
RETURNING id, created_at, title
`

func (q *Queries) CreateHashtag(ctx context.Context, title string) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, createHashtag, title)
	var i Hashtag
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const createHashtagPost = `-- name: CreateHashtagPost :one
INSERT INTO hashtag_post (post_id, hashtag_id)
VALUES ($1, $2)
RETURNING id, hashtag_id, post_id
`

type CreateHashtagPostParams struct {
	PostID    int32
	HashtagID int32
}

func (q *Queries) CreateHashtagPost(ctx context.Context, arg CreateHashtagPostParams) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, createHashtagPost, arg.PostID, arg.HashtagID)
	var i HashtagPost
	err := row.Scan(&i.ID, &i.HashtagID, &i.PostID)
	return i, err
}

const deleteHashtag = `-- name: DeleteHashtag :one
DELETE FROM hashtags
WHERE id = $1
RETURNING id, created_at, title
`

func (q *Queries) DeleteHashtag(ctx context.Context, id int32) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, deleteHashtag, id)
	var i Hashtag
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const deleteHashtagPost = `-- name: DeleteHashtagPost :one
DELETE FROM hashtag_post
WHERE id = $1
RETURNING id, hashtag_id, post_id
`

func (q *Queries) DeleteHashtagPost(ctx context.Context, id int32) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, deleteHashtagPost, id)
	var i HashtagPost
	err := row.Scan(&i.ID, &i.HashtagID, &i.PostID)
	return i, err
}

const getHashtagById = `-- name: GetHashtagById :one
SELECT id, created_at, title FROM hashtags
WHERE id = $1
`

func (q *Queries) GetHashtagById(ctx context.Context, id int32) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, getHashtagById, id)
	var i Hashtag
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const getHashtagByTitle = `-- name: GetHashtagByTitle :one
SELECT id, created_at, title FROM hashtags
WHERE title = $1
`

func (q *Queries) GetHashtagByTitle(ctx context.Context, title string) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, getHashtagByTitle, title)
	var i Hashtag
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const getHashtagPostById = `-- name: GetHashtagPostById :one
SELECT id, hashtag_id, post_id FROM hashtag_post
WHERE id = $1
`

func (q *Queries) GetHashtagPostById(ctx context.Context, id int32) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, getHashtagPostById, id)
	var i HashtagPost
	err := row.Scan(&i.ID, &i.HashtagID, &i.PostID)
	return i, err
}

const listHashtags = `-- name: ListHashtags :many
SELECT id, created_at, title FROM hashtags
ORDER BY id
`

func (q *Queries) ListHashtags(ctx context.Context) ([]Hashtag, error) {
	rows, err := q.db.QueryContext(ctx, listHashtags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hashtag
	for rows.Next() {
		var i Hashtag
		if err := rows.Scan(&i.ID, &i.CreatedAt, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHashtagsPost = `-- name: ListHashtagsPost :many
SELECT id, hashtag_id, post_id FROM hashtag_post
ORDER BY id
`

func (q *Queries) ListHashtagsPost(ctx context.Context) ([]HashtagPost, error) {
	rows, err := q.db.QueryContext(ctx, listHashtagsPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HashtagPost
	for rows.Next() {
		var i HashtagPost
		if err := rows.Scan(&i.ID, &i.HashtagID, &i.PostID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHashtag = `-- name: UpdateHashtag :one
UPDATE hashtags
SET title = $2
WHERE id = $1
RETURNING id, created_at, title
`

type UpdateHashtagParams struct {
	ID    int32
	Title string
}

func (q *Queries) UpdateHashtag(ctx context.Context, arg UpdateHashtagParams) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, updateHashtag, arg.ID, arg.Title)
	var i Hashtag
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}