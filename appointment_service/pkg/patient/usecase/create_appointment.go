package usecase

import (
	"context"
	"healthcare-service/domain"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
)

type IPatientUCase struct {
	PatientRepo repository.IPatientRepository
}

func NewPatientUsecase(patientRepo repository.IPatientRepository) usecase.IPatientUseCase {
	return IPatientUCase {
		PatientRepo: patientRepo,
	}
}

func (puc IPatientUCase) CreateAppointment(ctx context.Context, patient entity.Patient, doctor entity.Doctor) (string, error) {
	return domain.EmptyString, nil
}