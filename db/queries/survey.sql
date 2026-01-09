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

-- name: GetSurveyFullByID :many
SELECT
    s.id            AS survey_id,
    s.owner_id,
    s.title,
    s.description,
    s.public_id,
    s.is_active,

    q.id            AS question_id,
    q.text          AS question_text,
    q.type          AS question_type,
    q.is_required,
    q.rating_max,

    o.id            AS option_id,
    o.text          AS option_text

FROM surveys s
         LEFT JOIN survey_questions q
                   ON q.survey_id = s.id
         LEFT JOIN survey_question_options o
                   ON o.question_id = q.id
WHERE s.id = $1;

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