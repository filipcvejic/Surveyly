-- name: CreateAnswerChoice :one
INSERT INTO survey_answer_choices (
    id, answer_id, option_id
) VALUES (
    $1, $2 , $3
) RETURING *;

-- name: ListChoicesByAnswer :many
SELECT * FROM survey_answer_choices
WHERE answer_id = $1;

-- name: DeleteAnswerChoice :exec
DELETE FROM survey_answer_choices
WHERE answer_id = $1 AND option_id = $2;

-- name: DeleteAllAnswerChoices :exec
DELETE FROM survey_answer_choices
WHERE answer_id = $1;