// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (contents, user_id, post_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

type CreateCommentParams struct {
	Contents string
	UserID   int32
	PostID   int32
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.Contents, arg.UserID, arg.PostID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, deleteComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const getComment = `-- name: GetComment :one
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
WHERE id = $1
`

func (q *Queries) GetComment(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const listComments = `-- name: ListComments :many
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
ORDER BY id
`

func (q *Queries) ListComments(ctx context.Context) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Contents,
			&i.UserID,
			&i.PostID,
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

const listCommentsPost = `-- name: ListCommentsPost :many
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
WHERE post_id= $1
`

func (q *Queries) ListCommentsPost(ctx context.Context, postID int32) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listCommentsPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Contents,
			&i.UserID,
			&i.PostID,
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

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET contents = $2
WHERE id = $1
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

type UpdateCommentParams struct {
	ID       int32
	Contents string
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.Contents)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}
