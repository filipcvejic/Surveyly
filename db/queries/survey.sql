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

-- name: GetSurveyWithQuestionsAndOptions :many
SELECT
    s.id as survey_id,
    s.owner_id as survey_owner_id,
    s.title as survey_title,
    s.description as survey_description,
    s.public_id as survey_public_id,
    s.is_active as survey_is_active,
    s.created_at as survey_created_at,
    s.updated_at as survey_updated_at,
    sq.id as question_id,
    sq.text as question_text,
    sq.type as question_type,
    sq.is_required as question_is_required,
    sq.created_at as question_created_at,
    sqo.id as option_id,
    COALESCE(sqo.text, '') as option_text,
    sqo.created_at as option_created_at
FROM surveys s
LEFT JOIN survey_questions sq ON s.id = sq.survey_id
LEFT JOIN survey_question_options sqo ON sq.id = sqo.question_id
WHERE s.id = $1
ORDER BY sq.created_at, sqo.created_at;

-- name: ListSurveysByOwner :many
SELECT * FROM surveys
WHERE owner_id = $1
ORDER BY created_at DESC;

-- name: ListActiveSurveysByOwner :many
SELECT * FROM surveys
WHERE owner_id = $1 AND is_active = true
ORDER BY created_at DESC;

-- name: UpdateSurvey :exec
UPDATE surveys
SET
    title = $2,
    description = $3,
    updated_at = $4
WHERE id = $1;

-- name: DeleteSurvey :exec
DELETE FROM surveys
WHERE id = $1;