-- name: GetSurveyByID :one
SELECT * FROM surveys
WHERE id = $1 LIMIT 1;

-- name: GetSurveysByOwnerID :many
SELECT * FROM surveys
WHERE owner_id = $1;
    
-- name: CreateSurvey :one
INSERT INTO surveys (
    id, owner_id, title, description, public_id, is_active, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7          
)
RETURNING *;