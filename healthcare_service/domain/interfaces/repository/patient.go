package repository

import (
	"context"
	"healthcare-service/domain/entity"
)

type IPatientRepository interface {
	SavePatientDetails(ctx context.Context, patient entity.Patient) error
}