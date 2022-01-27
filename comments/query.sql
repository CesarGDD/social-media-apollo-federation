-- name: ListCommentsPost :many
SELECT * FROM comments
WHERE post_id= $1;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1;

-- name: ListComments :many
SELECT * FROM comments
ORDER BY id;

-- name: CreateComment :one
INSERT INTO comments (contents, user_id, post_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET contents = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING *;
