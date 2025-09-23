package services

import (
	"github.com/champion19/Flighthours_backend/core/domain"
	"github.com/champion19/Flighthours_backend/core/ports"
)

type service struct {
	repository ports.Repository
}


func NewService(repo ports.Repository) ports.Service {
	return &service{
		repository: repo,
	}
}

func (s service) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	return s.repository.GetEmployeeByEmail(email)
}

func (s service) RegisterEmployee(employee domain.Employee) (domain.Employee, error) {
	existingEmployee, err := s.GetEmployeeByEmail(employee.Email)
	if err == nil && existingEmployee != nil {
		return domain.Employee{}, domain.ErrDuplicateEmployee
	}

	employee.SetID()
	if err := employee.HashPassword(); err != nil {
		return domain.Employee{}, err
	}
	err = s.repository.Save(employee)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}
