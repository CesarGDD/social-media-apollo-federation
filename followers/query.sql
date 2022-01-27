-- name: GetFollower :one
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowers :many
SELECT * FROM followers
ORDER BY id;

-- name: ListFollowersById :many
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowersByLeaderId :many
SELECT * FROM followers
WHERE leader_id = $1;

-- name: CreateFollower :one
INSERT INTO followers (leader_id, follower_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING *;