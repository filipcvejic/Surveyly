-- +goose Up
-- +goose StatementBegin
CREATE TABLE survey_answer_choices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    answer_id UUID NOT NULL REFERENCES survey_answers(id) ON DELETE CASCADE,
    option_id UUID NOT NULL REFERENCES survey_question_options(id) ON DELETE CASCADE,
    
    UNIQUE(answer_id, option_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE survey_answer_choices
-- +goose StatementEnd
