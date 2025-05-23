package form

import (
	"amazing_form/internal/adapter/handler/dto/form_question"
	"encoding/json"
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormOutput struct {
	Id                 uint                               `json:"id"`
	CourseAssignmentId uint                               `json:"course_assignment_id"`
	MotherForm         *FormOutput                        `json:"mother_form"`
	FormQuestions      []form_question.FormQuestionOutput `json:"form_questions"`
}

func FormOutputFromDomain(f *domain.Form) *FormOutput {
	if f == nil {
		return nil
	}

	return &FormOutput{
		Id:                 f.ID,
		CourseAssignmentId: getUintOrZero(f.CourseAssignmentId),
		MotherForm:         FormOutputFromDomain(f.MotherForm),
		FormQuestions:      FormQuestionOutputsFromDomain(f.FormQuestions),
	}
}

func FormQuestionOutputsFromDomain(qs []domain.FormQuestion) []form_question.FormQuestionOutput {
	var result []form_question.FormQuestionOutput

	for _, q := range qs {
		var opts []string
		_ = json.Unmarshal([]byte(q.Options), &opts)
		result = append(result, form_question.FormQuestionOutput{
			ID:         q.ID,
			Question:   q.Question,
			Type:       string(rune(q.Type)),
			IsRequired: q.IsRequired,
			Options:    opts,
		})
	}
	return result
}

func getUintOrZero(ptr *uint) uint {
	if ptr == nil {
		return 0
	}
	return *ptr
}
