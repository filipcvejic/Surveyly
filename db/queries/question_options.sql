-- name: CreateQuestionOption :one
INSERT INTO survey_question_options (
    id, question_id, text,  created_at
) VALUES (
    $1, $2, $3, $4      
) RETURNING *;

-- name: CreateQuestionOptions :copyfrom
INSERT INTO survey_question_options (
    id, question_id, text, created_at
) VALUES (
    $1, $2, $3, $4
);

-- name: ListQuestionOptions :many
SELECT * FROM survey_question_options
WHERE question_id = $1;

-- name: UpdateQuestionOption :exec
UPDATE survey_question_options
SET
    text = $2
WHERE id = $1;

-- name: DeleteQuestionOption :exec
DELETE FROM survey_question_options
WHERE id = $1;

-- name: DeleteQuestionOptionsByQuestion :exec
DELETE FROM survey_question_options
WHERE question_id = $1;