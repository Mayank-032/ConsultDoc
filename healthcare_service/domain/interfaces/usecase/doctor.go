package usecase

import (
	"context"
)

type IDoctorUseCase interface {
	SendAppointmentLink(ctx context.Context, doctorId int, appointmentLink string) error
}