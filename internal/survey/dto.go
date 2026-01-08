package survey

import "github.com/google/uuid"

type CreateSurveyRequest struct {
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Questions   []CreateQuestionRequest `json:"questions" binding:"required,min=1,max=50"`
}

type CreateQuestionRequest struct {
	Text       string   `json:"text" validate:"required,min=3,max=500"`
	Type       string   `json:"type" validate:"required,oneof=text single_choice multiple_choice rating"`
	IsRequired bool     `json:"is_required"`
	Options    []string `json:"options"`
}

type SurveyResponse struct {
	ID          uuid.UUID          `json:"id"`
	OwnerID     uuid.UUID          `json:"owner_id"`
	Title       string             `json:"title"`
	Description string             `json:"description,omitempty"`
	PublicID    string             `json:"public_id"`
	IsActive    bool               `json:"is_active"`
	Questions   []QuestionResponse `json:"questions"`
}

type QuestionResponse struct {
	ID         uuid.UUID        `json:"id"`
	Text       string           `json:"text"`
	Type       string           `json:"type"`
	IsRequired bool             `json:"is_required"`
	Options    []OptionResponse `json:"options,omitempty"`
}

type OptionResponse struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
