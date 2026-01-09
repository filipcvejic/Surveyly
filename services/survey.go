package services

import (
	"context"
	"github.com/filipcvejic/surveyly/repositories"
)

type SurveyService struct {
	surveyRepo          *repositories.SurveyRepository
	surveyQuestionsRepo *repositories.SurveyQuestionsRepo
	surveyAnswersRepo   *repositories.SurveyAnswersRepo
}

func NewSurveyService(surveyRepo *repositories.SurveyRepository, surveyQuestionsRepo *repositories.SurveyQuestionsRepo, surveyAnswersRepo *repositories.SurveyAnswersRepo) {
	return &SurveyService{surveyRepo: surveyRepo, surveyQuestionsRepo: surveyQuestionsRepo, surveyAnswersRepo: surveyAnswersRepo}
}

func (s *SurveyService) CreateSurvey(ctx context.Context)
