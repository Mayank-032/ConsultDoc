package usecase

import (
	"context"
	"healthcare-service/domain/entity"
)

type ICommonUseCase interface {
	FetchDoctorsList(ctx context.Context, address entity.Address) ([]entity.Doctor, error)
}
