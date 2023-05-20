package repository

import (
	"context"
	"database/sql"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
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
	return nil
}