package form_question

type FormQuestionOutput struct {
	ID         uint     `json:"id"`
	Question   string   `json:"question"`
	Type       string   `json:"type"`
	IsRequired bool     `json:"is_required"`
	Options    []string `json:"options,omitempty"`
}
