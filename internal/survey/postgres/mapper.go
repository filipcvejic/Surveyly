package postgres

import (
	"github.com/filipcvejic/surveyly/db/sqlc"
	"github.com/filipcvejic/surveyly/internal/survey"
)

func mapSurvey(row sqlc.Survey) *survey.Survey {
	return &survey.Survey{
		ID:          row.ID,
		OwnerID:     row.OwnerID,
		Title:       row.Title,
		Description: row.Description,
		PublicID:    row.PublicID,
		IsActive:    row.IsActive,
	}
}
