package repository

import (
	"context"
	"healthcare-service/domain/entity"
)

type ICommonRepository interface {
	GetDoctorsList(ctx context.Context, address entity.Address) ([]entity.Doctor, error)
}
