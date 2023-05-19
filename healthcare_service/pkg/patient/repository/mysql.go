package repository

import (
	"context"
	"database/sql"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
)

type IPatientRepo struct {
	DB *sql.DB
}

func NewPatientRepo(db *sql.DB) repository.IPatientRepository {
	return IPatientRepo {
		DB: db,
	}
}

func (pr IPatientRepo) InsertAppointmentDetails(ctx context.Context, patient entity.Patient, doctor entity.Doctor) error {
	return nil
}