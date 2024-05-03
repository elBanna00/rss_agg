-- name: CreateFollowedFeed :one
INSERT INTO follow_feeds (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetFollowedFeed :many
SELECT * FROM follow_feeds WHERE user_id = $1;

-- name: DeleteFollowedFeed :exec
DELETE FROM follow_feeds WHERE  id = $1 AND user_id = $2;