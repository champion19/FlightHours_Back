package handlers


import (
	"errors"
	"net/http"

	domain "github.com/champion19/Flighthours_backend/core/domain"
	"github.com/gin-gonic/gin"
)

var (
	ErrUnmarshalBody  = errors.New("error unmarshal request body")
	ErrValidationUser = errors.New("error validation user")
)

type WebError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (h Handler) HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrUnmarshalBody):
		c.JSON(http.StatusBadRequest, WebError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	case errors.Is(err, ErrValidationUser):
		c.JSON(http.StatusBadRequest, WebError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrEmployeeCannotSave):
		c.JSON(http.StatusFailedDependency, WebError{
			Status:  http.StatusFailedDependency,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrDuplicateEmployee):
		c.JSON(http.StatusAlreadyReported, WebError{
			Status:  http.StatusAlreadyReported,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, WebError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrEmailAlreadyVerified):
		c.JSON(http.StatusConflict, WebError{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrEmployeeCannotGet):
		c.JSON(http.StatusNotFound, WebError{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrEmployeeCannotFound):
		c.JSON(http.StatusNotFound, WebError{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})

	default:
		c.JSON(http.StatusInternalServerError, WebError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
}
