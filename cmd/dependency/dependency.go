package dependency

import (


	"github.com/champion19/Flighthours_backend/platform/mysql"
	"github.com/champion19/Flighthours_backend/config"
)



type Dependencies struct {

	Config        *config.Config
}

func Init() (*Dependencies, error) {
	cfg := config.MustLoadConfig()

	_, err := mysql.GetDB(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config:        cfg,
	}, nil
}
