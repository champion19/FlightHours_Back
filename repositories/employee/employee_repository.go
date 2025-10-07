package employee

import (
	"database/sql"
	domain "github.com/champion19/Flighthours_backend/core/domain"
	"github.com/champion19/Flighthours_backend/core/ports"
	mysql"github.com/go-sql-driver/mysql"
)

const (
	querySave                 = "INSERT INTO employee(id,name,airline,email,password,email_confirmed,identification_number,bp,start_date,end_date,active) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	QueryByEmail              = "Select id,name,airline,email,password,email_confirmed,identification_number,bp,start_date,end_date,active FROM  employee WHERE email=?"
)

type repository struct {
	db         *sql.DB
	stmtSave   *sql.Stmt
	stmtGetByEmail *sql.Stmt
}

func NewRepository(db *sql.DB) (ports.Repository, error) {
	stmtSave, err := db.Prepare(querySave)
	if err != nil {
		return nil, domain.ErrUserCannotSave
	}
	stmtGetByEmail,err:=db.Prepare(QueryByEmail)
	if err != nil {
		return nil, domain.ErrUserCannotSave
	}
	return &repository{
		db: db,
		stmtSave: stmtSave,
		stmtGetByEmail: stmtGetByEmail,
	}, nil
}

func (r *repository) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	var e Employee
	err := r.stmtGetByEmail.QueryRow(email).Scan(
		&e.ID, &e.Name, &e.Airline, &e.Email, &e.Password, &e.Emailconfirmed, &e.IdentificationNumber, &e.Bp, &e.StartDate, &e.EndDate, &e.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPersonNotFound
		}
		return nil, err
	}
	d := e.ToDomain()
	return &d, nil
		}


func (r *repository) Save(employee domain.Employee) error {
	employeeToSave := Employee{
		ID:                   employee.ID,
		Name:                 employee.Name,
		Airline:              employee.Airline,
		Email:                employee.Email,
		Password:             employee.Password,
		Emailconfirmed:       employee.Emailconfirmed,
		IdentificationNumber: employee.IdentificationNumber,
		Bp:                   employee.Bp,
		StartDate:            employee.StartDate,
		EndDate:              employee.EndDate,
		Active:               employee.Active,
	}

	_, err := r.stmtSave.Exec(
		employeeToSave.ID,
		employeeToSave.Name,
		employeeToSave.Airline,
		employeeToSave.Email,
		employeeToSave.Password,
		employeeToSave.Emailconfirmed,
		employeeToSave.IdentificationNumber,
		employeeToSave.Bp,
		employeeToSave.StartDate,
		employeeToSave.EndDate,
		employeeToSave.Active,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return domain.ErrDuplicateUser
		} else {
			return domain.ErrUserCannotSave
		}
	}

	return nil
}

func (r *repository) Close() {
	if r.stmtSave != nil {
		r.stmtSave.Close()
	}
	if r.stmtGetByEmail != nil {
		r.stmtGetByEmail.Close()
	}
}


