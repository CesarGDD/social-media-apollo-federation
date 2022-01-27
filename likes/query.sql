-- name: ListLikesFromPost :many
SELECT * FROM post_likes
WHERE post_id= $1;

-- name: ListPostLikes :many
SELECT * FROM post_likes
ORDER BY id;

-- name: GetPostLike :one
SELECT * FROM post_likes
WHERE id = $1;

-- name: CreatePostLike :one
INSERT INTO post_likes (user_id, post_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePostLike :one
DELETE FROM post_likes
WHERE id = $1
RETURNING *;

-- name: ListCommentLikes :many
SELECT * FROM comment_likes
ORDER BY id;

-- name: ListLikesFromComment :many
SELECT * FROM comment_likes
WHERE comment_id = $1;

-- name: GetCommentLike :one
SELECT * FROM comment_likes
WHERE id = $1;

-- name: CreateCommentLike :one
INSERT INTO comment_likes (user_id, comment_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteCommentLike :one
DELETE FROM comment_likes
WHERE id = $1
RETURNING *;
