package form

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type CreateFormInput struct {
	MotherId      *uint               `form:"mother_form_id"`
	FormQuestions []FormQuestionInput `json:"form_questions" validate:"required,dive"`
}

func (i *CreateFormInput) ToDomain() *domain.Form {
	var questions []domain.FormQuestion
	for _, q := range i.FormQuestions {
		questions = append(questions, q.ToDomain())
	}
	return &domain.Form{
		MotherId:      i.MotherId,
		FormQuestions: questions,
	}
}
