package employee

import (
	"net/http"

	domain "github.com/champion19/Flighthours_backend/core/domain"
	"github.com/champion19/Flighthours_backend/core/ports"
	"github.com/champion19/Flighthours_backend/handlers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*handlers.Handler
}



func New(service ports.Service) *Handler {
	return &Handler{
		Handler: handlers.New(service),
	}
}

func (h Handler) GetEmployeeByEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")
		if email == "" {
			h.HandleError(c, domain.ErrEmployeeCannotFound)
			return
		}

		employee, err := h.EmployeeService.GetEmployeeByEmail(email)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := EmployeeResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h Handler) RegisterEmployee() func(c *gin.Context) {
	return func(c *gin.Context) {
		var employeeRequest EmployeeRequest
		err := c.ShouldBindJSON(&employeeRequest)
		if err != nil {
			h.HandleError(c, domain.ErrUnmarshalBody)
			return
		}

		err = employeeRequest.Validate()
		if err != nil {
			h.HandleError(c, domain.ErrValidationUser)
			return
		}

		domainEmployee, err := employeeRequest.ToDomain()
		if err != nil {
			h.HandleError(c, err)
			return
		}

		employee, err := h.EmployeeService.RegisterEmployee(domainEmployee)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := EmployeeResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Employee created successfully. Please check your email to verify your account.",
			"employee": response,
		})
	}
}
