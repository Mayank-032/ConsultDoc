package usecase

import (
	"context"
	"healthcare-service/domain/interfaces/usecase"
)

type DoctorUCase struct {
}

func NewDoctorUCase() usecase.IDoctorUseCase {
	return DoctorUCase{}
}

func (duc DoctorUCase) SendAppointmentLink(ctx context.Context, appointmentLink string) error {
	return nil
}