package repository

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-service/domain"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
	"log"
)

type DoctorRepo struct {
	DB *sql.DB
}

func NewDoctorRepo(db *sql.DB) repository.IDoctorRepository {
	return DoctorRepo {
		DB: db,
	}
}

func (dr DoctorRepo) FetchDoctorDetails(ctx context.Context, id int) (entity.Doctor, error) {
	conn, err := dr.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_db_connect\n\n", err.Error())
		return entity.Doctor{}, errors.New("unable to db connect")
	}
	defer conn.Close()

	var doctor entity.Doctor
	sqlQuery := "SELECT id, name, phone, ST_X(location), ST_Y(location), appointment_fees FROM doctors WHERE id = ? and status = ?"
	args := []interface{}{id, domain.StatusActive}
	err = dr.DB.QueryRowContext(ctx, sqlQuery, args...).Scan(&doctor)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_execute_query\n\n", err)
		return doctor, errors.New("unable to execute query")
	}
	return doctor, nil
}