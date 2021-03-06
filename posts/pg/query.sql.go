// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (url, caption, user_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, url, caption, user_id
`

type CreatePostParams struct {
	Url     string
	Caption string
	UserID  int32
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Url, arg.Caption, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.URL,
		&i.Caption,
		&i.UserID,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING id, created_at, updated_at, url, caption, user_id
`

func (q *Queries) DeletePost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.URL,
		&i.Caption,
		&i.UserID,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE id = $1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.URL,
		&i.Caption,
		&i.UserID,
	)
	return i, err
}

const getPostUserId = `-- name: GetPostUserId :one
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE user_id = $1
`

func (q *Queries) GetPostUserId(ctx context.Context, userID int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostUserId, userID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.URL,
		&i.Caption,
		&i.UserID,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
ORDER BY id
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.URL,
			&i.Caption,
			&i.UserID,
		); err != nil {
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

const listPostsById = `-- name: ListPostsById :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE id= $1
`

func (q *Queries) ListPostsById(ctx context.Context, id int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.URL,
			&i.Caption,
			&i.UserID,
		); err != nil {
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

const listPostsUser = `-- name: ListPostsUser :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE user_id= $1
`

func (q *Queries) ListPostsUser(ctx context.Context, userID int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.URL,
			&i.Caption,
			&i.UserID,
		); err != nil {
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET url = $2, caption = $3
WHERE id = $1
RETURNING id, created_at, updated_at, url, caption, user_id
`

type UpdatePostParams struct {
	ID      int32
	Url     string
	Caption string
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Url, arg.Caption)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.URL,
		&i.Caption,
		&i.UserID,
	)
	return i, err
}
