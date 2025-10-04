package ports

import "github.com/champion19/Flighthours_backend/core/domain"



type Service interface {
	RegisterEmployee(employee domain.Employee) (domain.Employee, error)
	GetEmployeeByEmail(email string) (*domain.Employee, error)
}
