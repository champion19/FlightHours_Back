package server

import (
	"log"
  "log/slog"
	"github.com/champion19/Flighthours_backend/cmd/dependency"
	"github.com/champion19/Flighthours_backend/handlers"
	"github.com/champion19/Flighthours_backend/middleware"
	"github.com/champion19/Flighthours_backend/platform/schema"

	"github.com/gin-gonic/gin"
)

func routing(app *gin.Engine, dependencies *dependency.Dependencies) {
	slog.Info("Setting up routes")

	handler := handlers.New(dependencies.EmployeeService)

	validators, err := schema.NewValidator(&schema.DefaultFileReader{})
	if err != nil {
		slog.Error("Error creating validator", slog.String("error", err.Error()))
		return
	}
	validator := middleware.NewMiddlewareValidator(validators)

	public := app.Group("/v1/flighthours")
	{
	public.POST("/register", validator.WithValidateRegister(),handler.RegisterEmployee())
	public.GET("/employee/:email", handler.GetEmployeeByEmail())
	}

}

func Bootstrap(app *gin.Engine) *dependency.Dependencies {
	dependencies, err := dependency.Init()
	if err != nil {
		log.Fatal("failed to init dependencies")
		return nil
	}
	routing(app, dependencies)
	return dependencies
}
