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
    text TEXT NOT NULL,
    type question_type NOT NULL,
    is_required BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE survey_questions;
DROP TYPE question_type;
-- +goose StatementEnd
