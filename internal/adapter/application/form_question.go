package application

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormQuestionUseCaseInterface interface {
	FindAll() ([]domain.FormQuestion, error)
	FindAllWithPagination(p int, l int) ([]domain.FormQuestion, error)
	FindByID(id uint) (*domain.FormQuestion, error)
	Create(q *domain.FormQuestion) error
	Update(q *domain.FormQuestion) error
	Delete(id uint) error
}

type FormQuestionRepositoryInterface interface {
	FindAll() ([]domain.FormQuestion, error)
	FindAllWithPagination(p int, l int) ([]domain.FormQuestion, error)
	FindByID(id uint) (*domain.FormQuestion, error)
	Create(formQuestion *domain.FormQuestion) error
	Update(formQuestion *domain.FormQuestion) error
	Delete(id uint) error
}

type FormQuestionUseCase struct {
	repo FormQuestionRepositoryInterface
}

func NewFormQuestionUseCase(r FormQuestionRepositoryInterface) FormQuestionUseCaseInterface {
	return &FormQuestionUseCase{repo: r}
}

func (uc *FormQuestionUseCase) FindAll() ([]domain.FormQuestion, error) {
	return uc.repo.FindAll()
}

func (uc *FormQuestionUseCase) FindAllWithPagination(page int, limit int) ([]domain.FormQuestion, error) {
	return uc.repo.FindAllWithPagination(page, limit)
}

func (uc *FormQuestionUseCase) FindByID(id uint) (*domain.FormQuestion, error) {
	return uc.repo.FindByID(id)
}

func (uc *FormQuestionUseCase) Create(q *domain.FormQuestion) error {
	return uc.repo.Create(q)
}

func (uc *FormQuestionUseCase) Update(q *domain.FormQuestion) error {
	return uc.repo.Update(q)
}

func (uc *FormQuestionUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
