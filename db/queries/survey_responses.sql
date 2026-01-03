-- name: CreateResponse :one
INSERT INTO survey_responses (
    id, survey_id, respondent_id, ip_address, created_at
) VALUES (
    $1, $2, $3, $4, $5          
) RETURNING *;

-- name: GetResponse :one
SELECT * FROM survey_responses
WHERE id = $1;

-- name: ListResponses :many
SELECT * FROM survey_responses
WHERE survey_id = $1
ORDER BY created_at DESC;