package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/filipcvejic/surveyly/db"
	"github.com/filipcvejic/surveyly/db/sqlc"
	"github.com/filipcvejic/surveyly/internal/survey"
	"github.com/google/uuid"
)

type surveyRepository struct {
	db *db.DB
}

func NewSurveyRepository(database *db.DB) survey.Repository {
	return &surveyRepository{
		db: database,
	}
}

func (r *surveyRepository) Create(ctx context.Context, s *survey.Survey) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := r.db.Query.WithTx(tx)

	if _, err := query.CreateSurvey(ctx, sqlc.CreateSurveyParams{
		ID:          s.ID,
		OwnerID:     s.OwnerID,
		Title:       s.Title,
		Description: s.Description,
		PublicID:    s.PublicID,
		IsActive:    s.IsActive,
	}); err != nil {
		return err
	}

	for _, question := range s.Questions {
		if _, err := query.CreateSurveyQuestion(ctx, sqlc.CreateSurveyQuestionParams{
			ID:         question.ID,
			SurveyID:   s.ID,
			Text:       question.Text,
			Type:       sqlc.QuestionType(question.Type),
			IsRequired: question.IsRequired,
		}); err != nil {
			return err
		}

		for _, option := range question.Options {
			if _, err := query.CreateQuestionOption(ctx, sqlc.CreateQuestionOptionParams{
				ID:         option.ID,
				QuestionID: question.ID,
				Text:       option.Text,
			}); err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

func (r *surveyRepository) FindByID(ctx context.Context, id uuid.UUID) (*survey.Survey, error) {
	row, err := r.db.Query.GetSurvey(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	survey := &survey.Survey{
		ID:          row.ID,
		OwnerID:     row.OwnerID,
		Title:       row.Title,
		Description: row.Description,
	}

	return mapSurvey(row), nil
}
