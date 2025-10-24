-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * from users
WHERE name = $1::TEXT;

-- name: Reset :exec
DELETE FROM users;

-- name: GetALLUsers :many
SELECT name from users;
