package repository

import (
	"amazing_form/internal/infrastructure/database"

	domain "github.com/MathieuRocher/amazing_domain"
	"gorm.io/gorm"
)

type FormQuestionType int

const (
	Field FormQuestionType = iota
	Rating
	Radio
	Select
)

var TypeName = map[FormQuestionType]string{
	Field:  "Field",
	Rating: "Rating",
	Radio:  "Radio",
	Select: "Select",
}

func (fqt FormQuestionType) String() string {
	return TypeName[fqt]
}

type FormQuestion struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	FormID     uint
	Question   string
	Type       FormQuestionType
	Options    string // Array JSON
	IsRequired bool
}

type FormQuestionRepository struct {
	db *gorm.DB
}

func NewFormQuestionRepository() *FormQuestionRepository {
	return &FormQuestionRepository{database.DB}
}

func (r *FormQuestionRepository) FindAll() ([]domain.FormQuestion, error) {
	var repoFormQuestions []FormQuestion
	if err := r.db.Find(&repoFormQuestions).Error; err != nil {
		return nil, err
	}

	var domainFormQuestions []domain.FormQuestion
	for _, repoFormQuestion := range repoFormQuestions {
		domainFormQuestion := repoFormQuestion.ToDomain()
		domainFormQuestions = append(domainFormQuestions, *domainFormQuestion)
	}

	return domainFormQuestions, nil
}

func (r *FormQuestionRepository) FindByID(id uint) (*domain.FormQuestion, error) {
	var obj FormQuestion
	if err := r.db.First(&obj, id).Error; err != nil {
		return nil, err
	}
	return obj.ToDomain(), nil
}

func (r *FormQuestionRepository) Create(obj *domain.FormQuestion) error {
	return r.db.Create(FormQuestionFromDomain(obj)).Error
}

func (r *FormQuestionRepository) Update(obj *domain.FormQuestion) error {
	return r.db.Save(FormQuestionFromDomain(obj)).Error
}

func (r *FormQuestionRepository) Delete(id uint) error {
	return r.db.Delete(&FormQuestion{}, id).Error
}

// ToDomain converts a repository FormQuestion to a domain FormQuestion
func (u *FormQuestion) ToDomain() *domain.FormQuestion {
	return &domain.FormQuestion{
		ID:         u.ID,
		FormID:     u.FormID,
		Question:   u.Question,
		Type:       domain.FormQuestionType(u.Type),
		Options:    u.Options,
		IsRequired: u.IsRequired,
	}
}

// FormQuestionFromDomain converts a domain FormQuestion to a repository FormQuestion
func FormQuestionFromDomain(u *domain.FormQuestion) *FormQuestion {
	return &FormQuestion{
		ID:         u.ID,
		FormID:     u.FormID,
		Question:   u.Question,
		Type:       FormQuestionType(u.Type),
		Options:    u.Options,
		IsRequired: u.IsRequired,
	}
}

func FormQuestionsFromDomain(assignments []domain.FormQuestion) []FormQuestion {
	var repoFormQuestions []FormQuestion
	for _, a := range assignments {
		repoFormQuestions = append(repoFormQuestions, *FormQuestionFromDomain(&a))
	}
	return repoFormQuestions
}

func ToDomainFormQuestions(repoQuestions []FormQuestion) []domain.FormQuestion {
	var domainQuestions []domain.FormQuestion
	for _, question := range repoQuestions {
		domainQuestions = append(domainQuestions, *question.ToDomain())
	}
	return domainQuestions
}
