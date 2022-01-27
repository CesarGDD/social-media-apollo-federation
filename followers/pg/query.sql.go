// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

const createFollower = `-- name: CreateFollower :one
INSERT INTO followers (leader_id, follower_id)
VALUES ($1, $2)
RETURNING id, created_at, leader_id, follower_id
`

type CreateFollowerParams struct {
	LeaderID   int32
	FollowerID int32
}

func (q *Queries) CreateFollower(ctx context.Context, arg CreateFollowerParams) (Follower, error) {
	row := q.db.QueryRowContext(ctx, createFollower, arg.LeaderID, arg.FollowerID)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.LeaderID,
		&i.FollowerID,
	)
	return i, err
}

const deleteFollower = `-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING id, created_at, leader_id, follower_id
`

func (q *Queries) DeleteFollower(ctx context.Context, id int32) (Follower, error) {
	row := q.db.QueryRowContext(ctx, deleteFollower, id)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.LeaderID,
		&i.FollowerID,
	)
	return i, err
}

const getFollower = `-- name: GetFollower :one
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE id = $1
`

func (q *Queries) GetFollower(ctx context.Context, id int32) (Follower, error) {
	row := q.db.QueryRowContext(ctx, getFollower, id)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.LeaderID,
		&i.FollowerID,
	)
	return i, err
}

const listFollowers = `-- name: ListFollowers :many
SELECT id, created_at, leader_id, follower_id FROM followers
ORDER BY id
`

func (q *Queries) ListFollowers(ctx context.Context) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.LeaderID,
			&i.FollowerID,
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

const listFollowersById = `-- name: ListFollowersById :many
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE id = $1
`

func (q *Queries) ListFollowersById(ctx context.Context, id int32) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowersById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.LeaderID,
			&i.FollowerID,
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

const listFollowersByLeaderId = `-- name: ListFollowersByLeaderId :many
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE leader_id = $1
`

func (q *Queries) ListFollowersByLeaderId(ctx context.Context, leaderID int32) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowersByLeaderId, leaderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.LeaderID,
			&i.FollowerID,
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