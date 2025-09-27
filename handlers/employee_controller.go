package handlers

import (
	"net/http"

	domain "github.com/champion19/Flighthours_backend/core/domain"
	"github.com/gin-gonic/gin"
)

func (h handler) GetEmployeeByEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")

		employee, err := h.EmployeeService.GetEmployeeByEmail(email)
		if err != nil {
			h.HandleError(c, err)
			return
		}
		c.JSON(http.StatusOK, employee)
	}
}

func (h handler) RegisterEmployee() func(c *gin.Context) {
	return func(c *gin.Context) {
		var employeeRequest EmployeeRequest
		if err := c.ShouldBindJSON(&employeeRequest); err != nil {
			h.HandleError(c, domain.ErrInvalidJSONFormat)
			return
		}

		employee, err := h.EmployeeService.RegisterEmployee(employeeRequest.ToDomain(),)
		if err != nil {

			switch err {
			case domain.ErrDuplicateUser:
				h.HandleError(c, domain.ErrDuplicateUser)
			case domain.ErrUserCannotSave:
				h.HandleError(c, domain.ErrUserCannotSave)
			default:
				h.HandleError(c, domain.ErrUserCannotSave)
			}
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

		c.JSON(http.StatusCreated, response)
	}
}
