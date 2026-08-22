-- name: GetUserByID :one
SELECT id, email, password_hash, created_at, updated_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ListUsers :many
SELECT id, email, created_at
FROM users
ORDER BY id;
