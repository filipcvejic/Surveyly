-- name: CreateSurveyQuestion :one
INSERT INTO survey_questions (
    id, survey_id, text, type, is_required, created_at               
) VALUES (
    $1, $2, $3, $4, $5, $6 
) RETURNING *;

-- name: GetSurveyQuestion :one
SELECT * FROM survey_questions
WHERE id = $1;

-- name: ListSurveyQuestions :many
SELECT * FROM survey_questions
WHERE survey_id = $1;

-- name: ListRequiredSurveyQuestions :many
SELECT * FROM survey_questions
WHERE survey_id = $1 AND is_required = true;

-- name: UpdateSurveyQuestion :exec
UPDATE survey_questions
SET
    text = $2,
    type = $3,
    is_required = $4
WHERE id = $1;

-- name: DeleteQuestion :exec
DELETE FROM survey_questions
WHERE id = $1;