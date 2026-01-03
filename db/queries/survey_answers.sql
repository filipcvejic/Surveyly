-- name: CreateAnswer :one
INSERT INTO survey_answers (
    id, response_id, question_id, answer_text, answer_numeric, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6          
) RETURNING *;

-- name: GetAnswer :one
SELECT * FROM survey_answers
WHERE id = $1;

-- name: ListAnswers :many
SELECT * FROM survey_answers
WHERE response_id = $1
ORDER BY created_at ASC;