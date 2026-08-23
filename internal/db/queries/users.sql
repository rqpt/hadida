-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, verification_token, verification_expires_at
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, email, is_verified, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, email, password_hash, is_verified, created_at, updated_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, is_verified, created_at, updated_at
FROM users
WHERE email = $1 LIMIT 1;

-- name: VerifyUserEmail :exec
UPDATE users
SET is_verified = TRUE,
    verification_token = NULL,
    verification_expires_at = NULL,
    updated_at = NOW()
WHERE verification_token = $1
  AND verification_expires_at > NOW();
