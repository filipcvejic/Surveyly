package survey

func ToSurveyResponse(s *Survey) SurveyResponse {
	questions := make([]QuestionResponse, 0, len(s.Questions))

	for _, q := range s.Questions {
		qr := QuestionResponse{
			ID:         q.ID,
			Text:       q.Text,
			Type:       string(q.Type),
			IsRequired: q.IsRequired,
		}

		if len(q.Options) > 0 {
			opts := make([]OptionResponse, 0, len(q.Options))
			for _, o := range q.Options {
				opts = append(opts, OptionResponse{
					ID:   o.ID,
					Text: o.Text,
				})
			}
			qr.Options = opts
		}

		questions = append(questions, qr)
	}

	return SurveyResponse{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		IsActive:    s.IsActive,
		Questions:   questions,
	}
}
