package dependency

import (

	mysql"github.com/champion19/Flighthours_backend/platform/mysql"
	"github.com/champion19/Flighthours_backend/config"
	"github.com/champion19/Flighthours_backend/core/ports"
	"github.com/champion19/Flighthours_backend/core/services"
	repo "github.com/champion19/Flighthours_backend/repositories/employee"

)



type Dependencies struct {
	EmployeeService ports.Service
	EmployeeRepository ports.Repository
	Config        *config.Config
}

func Init() (*Dependencies, error) {
	cfg := config.MustLoadConfig()

	db, err := mysql.GetDB(cfg.Database)
	if err != nil {
		return nil, err
	}
	employeeRepo:=repo.NewRepository(db)
	employeeService:=services.NewService(employeeRepo,cfg)

	return &Dependencies{
		EmployeeService: employeeService,
		EmployeeRepository: employeeRepo,
		Config:        cfg,
	}, nil
}
