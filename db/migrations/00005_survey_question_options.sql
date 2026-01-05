-- +goose Up
-- +goose StatementBegin
CREATE TABLE survey_question_options (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES survey_questions(id) ON DELETE CASCADE,
    option_text TEXT NOT NULL,
    index INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(question_id, position)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE survey_question_options
-- +goose StatementEnd
