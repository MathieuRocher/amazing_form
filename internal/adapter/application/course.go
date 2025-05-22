package application

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type CourseUsecaseInterface interface {
	FindAll() ([]domain.Course, error)
	FindByID(id uint) (*domain.Course, error)
	Create(user *domain.Course) error
	Update(user *domain.Course) error
	Delete(id uint) error
}

type CourseRepositoryInterface interface {
	FindAll() ([]domain.Course, error)
	FindByID(id uint) (*domain.Course, error)
	Create(course *domain.Course) error
	Update(course *domain.Course) error
	Delete(id uint) error
}

type CourseUsecase struct {
	courseRepository CourseRepositoryInterface
}

func NewCourseUsecase(courseRepository CourseRepositoryInterface) CourseUsecaseInterface {
	return &CourseUsecase{
		courseRepository: courseRepository,
	}
}

func (u *CourseUsecase) FindAll() ([]domain.Course, error) {
	return u.courseRepository.FindAll()
}

func (u *CourseUsecase) FindByID(id uint) (*domain.Course, error) {
	return u.courseRepository.FindByID(id)
}

func (u *CourseUsecase) Create(user *domain.Course) error {
	return u.courseRepository.Create(user)
}

func (u *CourseUsecase) Update(user *domain.Course) error {
	return u.courseRepository.Update(user)

}

func (u *CourseUsecase) Delete(id uint) error {
	return u.courseRepository.Delete(id)

}
