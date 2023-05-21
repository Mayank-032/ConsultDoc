package usecase

import (
	"context"
)

type IDoctorUseCase interface {
	SendAppointmentLink(ctx context.Context, appointmentLink string) error
}