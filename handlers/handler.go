package handlers

import "github.com/champion19/Flighthours_backend/core/ports"

type Handler struct {
	EmployeeService ports.Service
}

func New(service ports.Service) *Handler {
	return &Handler{
		EmployeeService: service,
	}
}
