package survey

import "errors"

var (
	ErrNoQuestions         = errors.New("survey must have at least one question")
	ErrInvalidQuestionType = errors.New("invalid question type")
	ErrOptionsRequired     = errors.New("options required for choice questions")
	ErrRatingMaxRequired   = errors.New("rating_max required for rating questions")
	ErrInvalidRatingMax    = errors.New("invalid rating_max value")
	ErrInvalidOwner        = errors.New("invalid owner")
	ErrSurveyNotFound      = errors.New("survey not found")
)
