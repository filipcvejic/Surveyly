-- +goose Up
-- +goose StatementBegin
CREATE TYPE question_type AS ENUM (
    'text',
    'single_choice',
    'multiple_choice',
    'rating'
);

CREATE TABLE survey_questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    survey_id UUID NOT NULL REFERENCES surveys(id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    question_type question_type NOT NULL,
    required BOOLEAN NOT NULL DEFAULT FALSE,
    position INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              
    UNIQUE(survey_id, position)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE survey_questions;
DROP TYPE question_type
-- +goose StatementEnd
