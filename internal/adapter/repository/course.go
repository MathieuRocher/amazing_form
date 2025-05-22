package repository

import (
	"amazing_form/internal/infrastructure/database"

	domain "github.com/MathieuRocher/amazing_domain"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Assignments []CourseAssignment
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{database.DB}
}

func (r *CourseRepository) FindAll() ([]domain.Course, error) {
	var repoCourses []Course
	if err := r.db.Find(&repoCourses).Error; err != nil {
		return nil, err
	}

	var domainCourses []domain.Course
	for _, repoCourse := range repoCourses {
		domainForm := repoCourse.ToDomain()
		domainCourses = append(domainCourses, *domainForm)
	}

	return domainCourses, nil
}

func (r *CourseRepository) FindAllWithPagination(page int, limit int) ([]domain.Course, error) {
	var repoCourses []Course

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	if err := r.db.
		Limit(limit).
		Offset(offset).
		Find(&repoCourses).Error; err != nil {
		return nil, err
	}

	var domainCourses []domain.Course
	for _, repoCourse := range repoCourses {
		domainCourse := repoCourse.ToDomain()
		domainCourses = append(domainCourses, *domainCourse)
	}

	return domainCourses, nil
}

func (r *CourseRepository) FindByID(id uint) (*domain.Course, error) {
	var obj Course
	if err := r.db.First(&obj, id).Error; err != nil {
		return nil, err
	}
	return obj.ToDomain(), nil
}

func (r *CourseRepository) Create(obj *domain.Course) error {
	return r.db.Create(CourseFromDomain(obj)).Error
}

func (r *CourseRepository) Update(obj *domain.Course) error {
	return r.db.Save(CourseFromDomain(obj)).Error
}

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&Course{}, id).Error
}

// ToDomain converts a repository Course to a domain Course
func (c *Course) ToDomain() *domain.Course {
	return &domain.Course{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Assignments: ToDomainCourseAssignments(c.Assignments),
	}
}

// CourseFromDomain converts a domain Course to a repository Course
func CourseFromDomain(c *domain.Course) *Course {
	return &Course{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Assignments: CourseAssignmentsFromDomain(c.Assignments),
	}
}
