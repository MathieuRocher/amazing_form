package repository

import (
	"amazing_form/internal/infrastructure/database"
	domain "github.com/MathieuRocher/amazing_domain"

	"gorm.io/gorm"
)

type Form struct {
	ID                 uint  `gorm:"primaryKey"`
	MotherId           *uint `gorm:"column:mother_id"`    // FK vers une autre Form
	MotherForm         *Form `gorm:"foreignKey:MotherId"` // relation explicite
	CourseAssignmentId *uint `gorm:"column:course_assignment_id"`
	CourseAssignment   *CourseAssignment
	FormQuestions      []FormQuestion `gorm:"foreignKey:FormID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type FormRepository struct {
	db *gorm.DB
}

func NewFormRepository() *FormRepository {
	return &FormRepository{database.DB}
}

func (r *FormRepository) FindAll() ([]domain.Form, error) {
	var repoForms []Form
	if err := r.db.Preload("FormQuestions").Find(&repoForms).Error; err != nil {
		return nil, err
	}

	var domainForms []domain.Form
	for _, repoForm := range repoForms {
		domainForm := repoForm.ToDomain()
		domainForms = append(domainForms, *domainForm)
	}

	return domainForms, nil
}

func (r *FormRepository) FindAllWithPagination(page, limit *int) ([]domain.Form, error) {
	var repoForms []Form

	offset := (*page - 1) * (*limit)
	if offset < 0 {
		offset = 0
	}

	// Récupération avec préchargement des relations et pagination
	if err := r.db.Preload("FormQuestions").
		Limit(*limit).
		Offset(offset).
		Find(&repoForms).Error; err != nil {
		return nil, err
	}

	var domainForms []domain.Form
	for _, repoForm := range repoForms {
		domainForm := repoForm.ToDomain()
		domainForms = append(domainForms, *domainForm)
	}

	return domainForms, nil

}

func (r *FormRepository) FindAllFiltered(courseID, classID, page, limit *int) ([]domain.Form, error) {
	var repoForms []Form

	db := r.db.Model(&Form{}).
		Preload("FormQuestions").
		Preload("MotherForm").
		Preload("CourseAssignment").
		Joins("LEFT JOIN course_assignments ON course_assignments.id = forms.course_assignment_id")

	if courseID != nil {
		db = db.Where("course_assignments.course_id = ?", *courseID)
	}
	if classID != nil {
		db = db.Where("course_assignments.class_group_id = ?", *classID)
	}
	if page != nil && limit != nil {
		offset := (*page - 1) * (*limit)
		db = db.Offset(offset).Limit(*limit)
	}

	if err := db.Find(&repoForms).Error; err != nil {
		return nil, err
	}

	var domainForms []domain.Form
	for _, f := range repoForms {
		domainForms = append(domainForms, *f.ToDomain())
	}
	return domainForms, nil
}

func (r *FormRepository) FindByID(id uint) (*domain.Form, error) {
	var obj Form
	if err := r.db.Preload("FormQuestions").First(&obj, id).Error; err != nil {
		return nil, err
	}
	return obj.ToDomain(), nil
}

func (r *FormRepository) Create(obj *domain.Form) error {
	return r.db.Create(FormFromDomain(obj)).Error
}

func (r *FormRepository) Update(obj *domain.Form) error {
	return r.db.Save(FormFromDomain(obj)).Error
}

func (r *FormRepository) Delete(id uint) error {
	return r.db.Delete(&Form{}, id).Error
}

// ToDomain converts a repository Form to a domain Form
func (fa *Form) ToDomain() *domain.Form {
	var mother *domain.Form
	if fa.MotherForm != nil {
		mother = fa.MotherForm.ToDomain()
	}
	var courseAssignment *domain.CourseAssignment
	if fa.CourseAssignment != nil {
		courseAssignment = fa.CourseAssignment.ToDomain()
	}

	return &domain.Form{
		ID:                 fa.ID,
		MotherId:           fa.MotherId,
		MotherForm:         mother,
		CourseAssignmentId: fa.CourseAssignmentId,
		CourseAssignment:   courseAssignment,
		FormQuestions:      ToDomainFormQuestions(fa.FormQuestions),
	}
}

// FormFromDomain converts a domain Form to a repository Form
func FormFromDomain(fa *domain.Form) *Form {
	if fa == nil {
		return nil
	}

	// Note : FormID sera rempli automatiquement par GORM après insertion
	form := &Form{
		ID:                 fa.ID,
		MotherId:           fa.MotherId,
		CourseAssignmentId: fa.CourseAssignmentId,
		FormQuestions:      []FormQuestion{},
	}

	for _, q := range fa.FormQuestions {
		form.FormQuestions = append(form.FormQuestions, FormQuestion{
			Question:   q.Question,
			Type:       FormQuestionType(q.Type),
			Options:    q.Options,
			IsRequired: q.IsRequired,
		})
	}

	return form
}
