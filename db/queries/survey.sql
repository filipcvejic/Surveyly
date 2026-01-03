-- name: CreateSurvey :one
INSERT INTO surveys (
    id, owner_id, title, description, public_id, is_active, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7          
)
RETURNING *;

-- name: GetSurvey :one
SELECT * FROM surveys
WHERE id = $1 LIMIT 1;

-- name: ListSurveysByOwner :many
SELECT * FROM surveys
WHERE owner_id = $1
ORDER BY created_at DESC;

-- name: ListActiveSurveysByOwner :many
SELECT * FROM surveys
WHERE owner_id = $1 AND is_active = true
ORDER BY created_at DESC;

-- name: DeleteSurvey :exec
DELETE FROM surveys
WHERE id = $1;
