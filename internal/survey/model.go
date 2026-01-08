package survey

import (
	"github.com/google/uuid"
	"time"
)

type Survey struct {
	ID          uuid.UUID
	OwnerID     uuid.UUID
	Title       string
	Description string
	PublicID    string
	IsActive    bool
	Questions   []Question
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Question struct {
	ID         uuid.UUID
	SurveyID   uuid.UUID
	Text       string
	Type       QuestionType
	Index      int
	IsRequired bool
	Options    []Option
	CreatedAt  time.Time
}

type Option struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Text       string
	Index      int
}

type QuestionType string

const (
	QuestionTypeText         QuestionType = "text"
	QuestionTypeSingleChoice QuestionType = "single_choice"
	QuestionTypeMultiChoice  QuestionType = "multi_choice"
	QuestionTypeRating       QuestionType = "rating"
)
