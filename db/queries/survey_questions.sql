-- name: CreateQuestion :one
INSERT INTO survey_questions (
    id, survey_id, question_text, question_type, required, position, created_at               
) VALUES (
    $1, $2, $3, $4, $5, $6, $7  
) RETURNING *;

-- name: GetQuestion :one
SELECT * FROM survey_questions
WHERE id = $1;

-- name: ListQuestions :many
SELECT * FROM survey_questions
WHERE survey_id = $1
ORDER BY position ASC;

-- name: ListRequiredQuestions :many
SELECT * FROM survey_questions
WHERE survey_id = $1 AND required = true
ORDER BY position ASC;

-- name: DeleteQuestion :exec
DELETE FROM survey_questions
WHERE id = $1;