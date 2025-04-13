-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, userid)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;