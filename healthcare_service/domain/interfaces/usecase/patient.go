package usecase

import (
	"context"
	"healthcare-service/domain/entity"
)

type IPatientUseCase interface {
	CreateAppointmentReceipt(ctx context.Context, patient entity.Patient, doctor entity.Doctor) error
}