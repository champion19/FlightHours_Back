package employee

import( "time"
"github.com/champion19/Flighthours_backend/core/domain"

)


type Employee struct {
	ID                   string    `db:"id"`
	Name                 string    `db:"name"`
	Airline              string    `db:"airline"`
	Email                string    `db:"email"`
	Password             string    `db:"password"`
	Emailconfirmed       bool      `db:"emailconfirmed"`
	IdentificationNumber string    `db:"identification_number"`
	Bp                   string    `db:"bp"`
	StartDate            time.Time `db:"start_date"`
	EndDate              time.Time `db:"end_date"`
	Active               bool      `db:"active"`
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

func FromDomain(d domain.Employee) Employee {
	return Employee{
		ID:                   d.ID,
		Name:                 d.Name,
		Email:                d.Email,
		Airline:              d.Airline,
		Password:             d.Password,
		Emailconfirmed:       d.Emailconfirmed,
		IdentificationNumber: d.IdentificationNumber,
		Bp:                   d.Bp,
		StartDate:            d.StartDate,
		EndDate:              d.EndDate,
		Active:               d.Active,
	}
}
