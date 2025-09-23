package server

import (
	"log"
  "log/slog"
	"github.com/champion19/Flighthours_backend/cmd/dependency"
	"github.com/champion19/Flighthours_backend/handlers/employee"

	"github.com/gin-gonic/gin"
)

func routing(app *gin.Engine, dependencies *dependency.Dependencies) {
	slog.Info("Setting up routes")

	handler := employee.New(dependencies.EmployeeService)

	app.POST("/v1/employees", handler.RegisterEmployee())

}

func Boostrap(app *gin.Engine) *dependency.Dependencies {
	dependencies, err := dependency.Init()
	if err != nil {
		log.Fatal("failed to init dependencies")
		return nil
	}
	routing(app, dependencies)
	return dependencies
}
