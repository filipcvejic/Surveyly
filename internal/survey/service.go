package survey

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewSurveyService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateSurveyWithQuestions(
	ctx context.Context,
	ownerID uuid.UUID,
	req CreateSurveyRequest,
) (*Survey, error) {

	if ownerID == uuid.Nil {
		return nil, ErrInvalidOwner
	}

	if len(req.Questions) == 0 {
		return nil, ErrNoQuestions
	}

	survey := &Survey{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Title:       req.Title,
		Description: req.Description,
		IsActive:    false,
		PublicID:    uuid.NewString(),
	}

	questions := make([]Question, 0, len(req.Questions))

	for _, q := range req.Questions {
		question := Question{
			ID:         uuid.New(),
			Text:       q.Text,
			Type:       QuestionType(q.Type),
			IsRequired: q.IsRequired,
		}

		switch question.Type {

		case QuestionTypeSingleChoice, QuestionTypeMultiChoice:
			if len(q.Options) == 0 {
				return nil, ErrOptionsRequired
			}

			options := make([]Option, 0, len(q.Options))
			for _, opt := range q.Options {
				options = append(options, Option{
					ID:   uuid.New(),
					Text: opt,
				})
			}

			question.Options = options
		case QuestionTypeText:
		case QuestionTypeRating:
		default:
			return nil, ErrInvalidQuestionType
		}

		questions = append(questions, question)
	}

	survey.Questions = questions

	if err := s.repo.Create(ctx, survey); err != nil {
		return nil, err
	}

	return survey, nil
}

func (s *Service) GetSurveyByID(ctx context.Context, id uuid.UUID) (*Survey, error) {
	if id == uuid.Nil {
		return nil, ErrSurveyNotFound
	}

	survey, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if survey == nil {
		return nil, ErrSurveyNotFound
	}

	return survey, nil
}
