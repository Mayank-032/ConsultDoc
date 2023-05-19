package usecase

import (
	"context"
	"healthcare-service/domain/entity"
)

type IPatientUseCase interface {
	CreateAppointment(ctx context.Context, patient entity.Patient, doctor entity.Doctor) (string, error)
}