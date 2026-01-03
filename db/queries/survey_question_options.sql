-- name: CreateOption :one
INSERT INTO survey_question_options (
    id, question_id, option_text, position, created_at
) VALUES (
    $1, $2, $3, $4, $5          
) RETURNING *;

-- name: CreateOptions :copyfrom
INSERT INTO survey_question_options (
    id, question_id, option_text, position, created_at
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: ListOptions :many
SELECT * FROM survey_question_options
WHERE question_id = $1
ORDER BY position ASC;