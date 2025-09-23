package employee

import( "time"
"github.com/champion19/Flighthours_backend/core/domain"

)


type Employee struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Airline              string    `json:"airline"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Emailconfirmed       bool      `json:"emailconfirmed"`
	IdentificationNumber string    `json:"identification_number"`
	Bp                   string    `json:"bp"`
	StartDate            time.Time `json:"start_date"`
	EndDate              time.Time `json:"end_date"`
	Active               bool      `json:"active"`
}

func (e Employee) ToDomain() domain.Employee {
	return domain.Employee{
		ID:                   e.ID,
		Name:                 e.Name,
		Email:                e.Email,
		Airline:              e.Airline,
		Password:             e.Password,
		Emailconfirmed:       e.Emailconfirmed,
		IdentificationNumber: e.IdentificationNumber,
		Bp:                   e.Bp,
		StartDate:            e.StartDate,
		EndDate:              e.EndDate,
		Active:               e.Active,
	}
}
