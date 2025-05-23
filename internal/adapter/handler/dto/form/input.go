package form

import (
	"amazing_form/internal/adapter/handler/dto/form_question"
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormInput struct {
	MotherId           *uint                             `form:"mother_form_id"`
	CourseAssignmentId *uint                             `form:"course_assignment_id"`
	FormQuestions      []form_question.FormQuestionInput `json:"form_questions" validate:"required,dive"`
}

func (i *FormInput) ToDomain() *domain.Form {
	var questions []domain.FormQuestion
	for _, q := range i.FormQuestions {
		questions = append(questions, q.ToDomain())
	}
	return &domain.Form{
		MotherId:           i.MotherId,
		CourseAssignmentId: i.CourseAssignmentId,
		FormQuestions:      questions,
	}
}
