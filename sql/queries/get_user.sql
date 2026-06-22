-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: GetUsers :many
SELECT * FROM users;