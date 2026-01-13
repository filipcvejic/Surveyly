package survey

import (
	"encoding/json"
	"github.com/filipcvejic/surveyly/internal/httpx"
	"github.com/filipcvejic/surveyly/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service, validator: validator.New()}
}

func (h *Handler) CreateSurvey(w http.ResponseWriter, r *http.Request) {
	var req CreateSurveyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ownerID, ok := middleware.GetUserID(r)
	if !ok || ownerID == uuid.Nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	survey, err := h.service.CreateSurveyWithQuestions(r.Context(), ownerID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.WriteJSON(w, 201, ToSurveyResponse(survey))
}

func (h *Handler) GetSurveyByID(w http.ResponseWriter, r *http.Request) {
	surveyIDStr := chi.URLParam(r, "id")

	surveyID, err := uuid.Parse(surveyIDStr)
	if err != nil {
		http.Error(w, "invalid survey id", 400)
		return
	}

	survey, err := h.service.GetSurveyDetails(r.Context(), surveyID)
	if err != nil {
		http.Error(w, err.Error(), 400) // Return actual error to see what's wrong
		return
	}

	httpx.WriteJSON(w, http.StatusOK, ToSurveyResponse(survey))
}
