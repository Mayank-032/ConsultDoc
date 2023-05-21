package usecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-service/domain"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	"log"
	"net/smtp"
)

type DoctorUCase struct {
	SMTPCred   domain.SMTP_Cred
	DoctorRepo repository.IDoctorRepository
}

func NewDoctorUCase(smtpCred domain.SMTP_Cred, doctorRepo repository.IDoctorRepository) usecase.IDoctorUseCase {
	return DoctorUCase{
		SMTPCred:   smtpCred,
		DoctorRepo: doctorRepo,
	}
}

func (duc DoctorUCase) SendAppointmentLink(ctx context.Context, doctorId int, appointmentLink string) error {
	doctor, err := duc.DoctorRepo.FetchDoctorDetails(ctx, doctorId)
	if err != nil {
		log.Printf("Error: %v,\n unable to fetch doctor details\n\n", err.Error())
		return errors.New("unable to fetch doctor details")
	}

	address := fmt.Sprintf("%v:%v", duc.SMTPCred.SMTPHost, duc.SMTPCred.SMTPPort)
	mailSubject := "Regarding Patient's Appointment\n\n"
	mailMessage := fmt.Sprintf("Dear %v,\n\nPlease find attached patient's appointment: %v\n\nWarm Regards,\nConsultDoc", doctor.Name, appointmentLink)
	mailBody := []byte(mailSubject + mailMessage)

	auth := smtp.PlainAuth("", duc.SMTPCred.BusinessEmail, duc.SMTPCred.BusinessPassword, duc.SMTPCred.SMTPHost)
	err = smtp.SendMail(address, auth, duc.SMTPCred.BusinessEmail, []string{doctor.Email}, mailBody)
	if err != nil {
		log.Printf("Error: %v,\n unable to mail doctor\n\n", err.Error())
		return errors.New("unable to mail doctor")
	}
	return nil
}
