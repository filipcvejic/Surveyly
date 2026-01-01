-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    id, email, username, password_hash, created_at
) VALUES (
    $1, $2,  $3, $4, $5
)
RETURNING *;