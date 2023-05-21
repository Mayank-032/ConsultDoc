package repository

import (
	"context"
	"healthcare-service/domain/entity"
)

type IDoctorRepository interface {
	FetchDoctorDetails(ctx context.Context, id int) (entity.Doctor, error)
}