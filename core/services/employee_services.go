package services

import (
	"github.com/champion19/Flighthours_backend/core/domain"
	"github.com/champion19/Flighthours_backend/core/ports"
	"github.com/champion19/Flighthours_backend/config"

)

type service struct {
	repository ports.Repository

	config *config.Config
}


func NewService(repo ports.Repository,cfg *config.Config) ports.Service {
	return &service{
		repository: repo,
		config: cfg,
	}
}

func (s service) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	return s.repository.GetEmployeeByEmail(email)
}

func (s service) RegisterEmployee(employee domain.Employee) (domain.Employee, error) {
	existingEmployee, err := s.repository.GetEmployeeByEmail(employee.Email)
	if err == nil && existingEmployee != nil {
		return domain.Employee{}, domain.ErrDuplicateUser
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
