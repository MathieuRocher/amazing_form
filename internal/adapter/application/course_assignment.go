package application

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type CourseAssignmentUseCaseInterface interface {
	FindAll() ([]domain.CourseAssignment, error)
	FindAllWithPagination(p int, l int) ([]domain.CourseAssignment, error)
	FindByID(id uint) (*domain.CourseAssignment, error)
	Create(q *domain.CourseAssignment) error
	Update(q *domain.CourseAssignment) error
	Delete(id uint) error
}

type CourseAssignmentRepositoryInterface interface {
	FindAll() ([]domain.CourseAssignment, error)
	FindAllWithPagination(p int, l int) ([]domain.CourseAssignment, error)
	FindByID(id uint) (*domain.CourseAssignment, error)
	Create(courseAssignment *domain.CourseAssignment) error
	Update(courseAssignment *domain.CourseAssignment) error
	Delete(id uint) error
}

type CourseAssignmentUseCase struct {
	repo CourseAssignmentRepositoryInterface
}

func NewCourseAssignmentUseCase(r CourseAssignmentRepositoryInterface) CourseAssignmentUseCaseInterface {
	return &CourseAssignmentUseCase{repo: r}
}

func (uc *CourseAssignmentUseCase) FindAll() ([]domain.CourseAssignment, error) {
	return uc.repo.FindAll()
}

func (uc *CourseAssignmentUseCase) FindAllWithPagination(page int, limit int) ([]domain.CourseAssignment, error) {
	return uc.repo.FindAllWithPagination(page, limit)
}

func (uc *CourseAssignmentUseCase) FindByID(id uint) (*domain.CourseAssignment, error) {
	return uc.repo.FindByID(id)
}

func (uc *CourseAssignmentUseCase) Create(q *domain.CourseAssignment) error {
	return uc.repo.Create(q)
}

func (uc *CourseAssignmentUseCase) Update(q *domain.CourseAssignment) error {
	return uc.repo.Update(q)
}

func (uc *CourseAssignmentUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
