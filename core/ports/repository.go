package ports

import "github.com/champion19/Flighthours_backend/core/domain"

type Repository interface {
	Save(employee domain.Employee) error
	GetEmployeeByEmail(email string) (*domain.Employee, error)
}
