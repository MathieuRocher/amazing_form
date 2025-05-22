package form

import (
	domain "github.com/MathieuRocher/amazing_domain"

	"strings"
)

type FormQuestionInput struct {
	Question   string `json:"question" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Options    string `json:"options,omitempty"`
	IsRequired bool   `json:"is_required"`
}

func (i *FormQuestionInput) ToDomain() domain.FormQuestion {
	return domain.FormQuestion{
		Question:   i.Question,
		Type:       parseFormQuestionType(i.Type),
		Options:    i.Options,
		IsRequired: i.IsRequired,
	}
}

func parseFormQuestionType(input string) domain.FormQuestionType {
	switch strings.ToLower(input) {
	case "field":
		return domain.Field
	case "rating":
		return domain.Rating
	case "radio":
		return domain.Radio
	case "select":
		return domain.Select
	default:
		return domain.Field
	}
}
