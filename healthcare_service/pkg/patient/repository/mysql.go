package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
	"log"
)

type PatientRepo struct {
	DB *sql.DB
}

func NewPatientRepo(db *sql.DB) repository.IPatientRepository {
	return PatientRepo {
		DB: db,
	}
}

func (pr PatientRepo) SavePatientDetails(ctx context.Context, patient entity.Patient) error {
	conn, err := pr.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_db_connect\n\n", err.Error())
		return errors.New("unable to db connect")
	}
	defer conn.Close()

	sqlQuery := "INSERT INTO healthcareServiceDB.patient_info(name, phone, location, appointment_receipt) VALUES (?, ?, ST_GeomFromText(?), ?);"
	args := []interface{}{patient.Name, patient.Phone, fmt.Sprintf("Point(%v %v)", patient.Address.Latitude, patient.Address.Longitude), patient.AppointmentLink}
	_, err = conn.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_execute_sql_query\n\n", err.Error())
		return errors.New("unable to execute sql query")
	}
	return nil
}