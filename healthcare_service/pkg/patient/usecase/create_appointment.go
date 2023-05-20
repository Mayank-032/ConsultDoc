package usecase

import (
	"context"
	"healthcare-service/domain"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/usecase"

	"cloud.google.com/go/storage"
)

type PatientUCase struct {
	Client     *storage.Client
}

func NewPatientUCase(client *storage.Client) usecase.IPatientUseCase {
	return PatientUCase{
		Client: client,
	}
}

func (puc PatientUCase) CreateAppointmentReceipt(ctx context.Context, patient entity.Patient, doctor entity.Doctor) (string, error) {
	return domain.EmptyString, nil
}