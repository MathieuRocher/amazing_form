package repository

import (
	"amazing_form/internal/infrastructure/database"

	domain "github.com/MathieuRocher/amazing_domain"

	"gorm.io/gorm"
)

type CourseAssignment struct {
	ID           uint `gorm:"primaryKey"`
	CourseID     uint
	ClassGroupID uint
	TrainerID    uint

	Course Course
	Forms  []Form `gorm:"foreignKey:CourseAssignmentId"`
}

type CourseAssignmentRepository struct {
	db *gorm.DB
}

func NewCourseAssignmentRepository() *CourseAssignmentRepository {
	return &CourseAssignmentRepository{database.DB}
}

func (r *CourseAssignmentRepository) FindAll() ([]domain.CourseAssignment, error) {
	var repoCourseAssignments []CourseAssignment
	if err := r.db.Find(&repoCourseAssignments).Error; err != nil {
		return nil, err
	}

	var domainCourseAssignments []domain.CourseAssignment
	for _, repoCourseAssignment := range repoCourseAssignments {
		domainCourseAssignment := repoCourseAssignment.ToDomain()
		domainCourseAssignments = append(domainCourseAssignments, *domainCourseAssignment)
	}

	return domainCourseAssignments, nil
}

func (r *CourseAssignmentRepository) FindAllWithPagination(page int, limit int) ([]domain.CourseAssignment, error) {
	var repoCourseAssignments []CourseAssignment

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	if err := r.db.
		Limit(limit).
		Offset(offset).
		Find(&repoCourseAssignments).Error; err != nil {
		return nil, err
	}

	var domainCourseAssignments []domain.CourseAssignment
	for _, repoCourseAssignment := range repoCourseAssignments {
		domainCourseAssignment := repoCourseAssignment.ToDomain()
		domainCourseAssignments = append(domainCourseAssignments, *domainCourseAssignment)
	}

	return domainCourseAssignments, nil
}

func (r *CourseAssignmentRepository) FindByID(id uint) (*domain.CourseAssignment, error) {
	var obj CourseAssignment
	if err := r.db.First(&obj, id).Error; err != nil {
		return nil, err
	}
	return obj.ToDomain(), nil
}

func (r *CourseAssignmentRepository) Create(obj *domain.CourseAssignment) error {
	return r.db.Create(CourseAssignmentFromDomain(obj)).Error
}

func (r *CourseAssignmentRepository) Update(obj *domain.CourseAssignment) error {
	return r.db.Save(CourseAssignmentFromDomain(obj)).Error
}

func (r *CourseAssignmentRepository) Delete(id uint) error {
	return r.db.Delete(&CourseAssignment{}, id).Error
}

// ToDomain converts a repository CourseAssignment to a domain CourseAssignment
func (ca *CourseAssignment) ToDomain() *domain.CourseAssignment {
	return &domain.CourseAssignment{
		ID:           ca.ID,
		CourseID:     ca.CourseID,
		ClassGroupID: ca.ClassGroupID,
		TrainerID:    ca.TrainerID,
		Course:       *ca.Course.ToDomain(),
	}
}

// CourseAssignmentFromDomain converts a domain CourseAssignment to a repository CourseAssignment
func CourseAssignmentFromDomain(ca *domain.CourseAssignment) *CourseAssignment {
	return &CourseAssignment{
		ID:           ca.ID,
		CourseID:     ca.CourseID,
		ClassGroupID: ca.ClassGroupID,
		TrainerID:    ca.TrainerID,
		Course:       *CourseFromDomain(&ca.Course),
	}
}

func CourseAssignmentsFromDomain(assignments []domain.CourseAssignment) []CourseAssignment {
	var repoAssignments []CourseAssignment
	for _, a := range assignments {
		repoAssignments = append(repoAssignments, *CourseAssignmentFromDomain(&a))
	}
	return repoAssignments
}

func ToDomainCourseAssignments(repoCourseAssignments []CourseAssignment) []domain.CourseAssignment {
	var domainCourseAssignments []domain.CourseAssignment
	for _, ca := range repoCourseAssignments {
		domainCourseAssignments = append(domainCourseAssignments, *ca.ToDomain())
	}
	return domainCourseAssignments
}
