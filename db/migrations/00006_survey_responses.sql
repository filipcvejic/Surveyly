-- +goose Up
-- +goose StatementBegin
CREATE TABLE survey_responses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    survey_id UUID NOT NULL REFERENCES surveys(id) ON DELETE CASCADE,
    respondent_id UUID,
    ip_address INET,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(survey_id, ip_address)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE survey_responses
-- +goose StatementEnd
