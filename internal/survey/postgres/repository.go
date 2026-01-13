package postgres

import (
	"context"
	"database/sql"
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

func (r *surveyRepository) GetByIDWithDetails(ctx context.Context, surveyID uuid.UUID) (*survey.Survey, error) {
	rows, err := r.db.Query.GetSurveyWithQuestionsAndOptions(ctx, surveyID)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, sql.ErrNoRows
	}

	firstRow := rows[0]
	result := &survey.Survey{
		ID:          firstRow.SurveyID,
		OwnerID:     firstRow.SurveyOwnerID,
		Title:       firstRow.SurveyTitle,
		Description: firstRow.SurveyDescription,
		PublicID:    firstRow.SurveyPublicID,
		IsActive:    firstRow.SurveyIsActive,
		Questions:   []survey.Question{},
		CreatedAt:   firstRow.SurveyCreatedAt,
		UpdatedAt:   firstRow.SurveyUpdatedAt,
	}

	questionMap := make(map[uuid.UUID]*survey.Question)

	for _, row := range rows {
		if !row.QuestionID.Valid {
			continue
		}

		questionID := row.QuestionID.Bytes

		if _, exists := questionMap[questionID]; !exists {
			questionMap[questionID] = &survey.Question{
				ID:         questionID,
				SurveyID:   surveyID,
				Text:       row.QuestionText,
				Type:       survey.QuestionType(row.QuestionType.QuestionType),
				IsRequired: row.QuestionIsRequired.Bool,
				CreatedAt:  row.QuestionCreatedAt.Time,
				Options:    []survey.Option{},
			}
		}

		if row.OptionID.Valid {
			questionMap[questionID].Options = append(
				questionMap[questionID].Options,
				survey.Option{
					ID:         row.OptionID.Bytes,
					QuestionID: questionID,
					Text:       row.OptionText,
				},
			)
		}
	}

	questions := make([]survey.Question, 0, len(questionMap))
	for _, q := range questionMap {
		questions = append(questions, *q)
	}

	result.Questions = questions

	return result, nil
}
