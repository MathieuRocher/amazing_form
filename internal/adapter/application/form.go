package application

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormUseCaseInterface interface {
	FindAll() ([]domain.Form, error)
	FindByID(id uint) (*domain.Form, error)
	Create(q *domain.Form) error
	Update(q *domain.Form) error
	Delete(id uint) error
}

type FormRepositoryInterface interface {
	FindAll() ([]domain.Form, error)
	FindByID(id uint) (*domain.Form, error)
	Create(form *domain.Form) error
	Update(form *domain.Form) error
	Delete(id uint) error
}

type FormUseCase struct {
	repo FormRepositoryInterface
}

func NewFormUseCase(r FormRepositoryInterface) FormUseCaseInterface {
	return &FormUseCase{repo: r}
}

func (uc *FormUseCase) FindAll() ([]domain.Form, error) {
	return uc.repo.FindAll()
}

func (uc *FormUseCase) FindByID(id uint) (*domain.Form, error) {
	return uc.repo.FindByID(id)
}

func (uc *FormUseCase) Create(q *domain.Form) error {
	return uc.repo.Create(q)
}

func (uc *FormUseCase) Update(q *domain.Form) error {
	return uc.repo.Update(q)
}

func (uc *FormUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
