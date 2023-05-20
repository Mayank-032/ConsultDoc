package usecase

import (
	"context"
	"errors"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	"log"
)

type CommonUCase struct {
	CommonRepo repository.ICommonRepository
}

func NewCommonUCase(commonRepo repository.ICommonRepository) usecase.ICommonUseCase {
	return CommonUCase {
		CommonRepo: commonRepo,
	}
}

func (cuc CommonUCase) FetchDoctorsList(ctx context.Context, address entity.Address) ([]entity.Doctor, error) {
	doctors, err := cuc.CommonRepo.GetDoctorsList(ctx, address)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_fetch_doctors_from_database\n\n", err.Error())
		return nil, errors.New("unable to fetch doctors")
	}

	for _, doctor := range doctors {
		slots, err := cuc.CommonRepo.GetDoctorSlots(ctx, doctor.Id)
		if err != nil {
			log.Printf("Error: %v\n, unable_to_fetch_doctor_slots\n\n", err.Error())
			// assigning an empty array so that operation continues for other doctors
			doctor.Slots = make([]entity.Slot, 0)
		} else {
			doctor.Slots = slots
		}
	}
	return doctors, nil
}