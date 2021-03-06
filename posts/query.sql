-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetPostUserId :one
SELECT * FROM posts
WHERE user_id = $1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id;

-- name: ListPostsById :many
SELECT * FROM posts
WHERE id= $1;

-- name: ListPostsUser :many
SELECT * FROM posts
WHERE user_id= $1;

-- name: CreatePost :one
INSERT INTO posts (url, caption, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET url = $2, caption = $3
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING *;