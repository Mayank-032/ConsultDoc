package repository

import (
	"healthcare-service/domain/entity"
	"context"
)

type IPatientRepository interface {
	InsertAppointmentDetails(ctx context.Context, patient entity.Patient, doctor entity.Doctor) error
}