package usecase

import (
	"context"
	"fmt"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/usecase"
	"log"

	"cloud.google.com/go/storage"
	"github.com/jung-kurt/gofpdf"
)

type PatientUCase struct {
	Client     *storage.Client
}

func NewPatientUCase(client *storage.Client) usecase.IPatientUseCase {
	return PatientUCase{
		Client: client,
	}
}

func (puc PatientUCase) CreateAppointmentReceipt(ctx context.Context, patient entity.Patient, doctor entity.Doctor) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// set page margin
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()
	pdf.SetMargins(leftMargin, topMargin, rightMargin)

	// add Patient Name
	pdf.CellFormat(0, 10, fmt.Sprintf("%v", patient.Name), "", 1, "C", false, 0, "")

	// add Patient Details
	pdf.CellFormat(0, 10, "Patient Details", "", 1, "L", false, 0, "")
	pdf.MultiCell(0, 10, fmt.Sprintf("%v", patient.Phone), "", "L", false)
	pdf.MultiCell(0, 10, fmt.Sprintf("%v", patient.Address), "", "L", false)

	// add Doctor Details
	pdf.CellFormat(0, 10, "Doctor Details", "", 1, "L", false, 0, "")
	pdf.MultiCell(0, 10, fmt.Sprintf("%v", doctor.Phone), "", "L", false)
	pdf.MultiCell(0, 10, fmt.Sprintf("Latitude: %v, Longitude: %v", doctor.Address.Latitude, doctor.Address.Latitude), "", "L", false)
	pdf.MultiCell(0, 10, fmt.Sprintf("Start-Time: %v, End-Time: %v", doctor.Slots[0].StartTime, doctor.Slots[0].EndTime), "", "L", false)

	// add fee details
	pdf.CellFormat(0, 10, "Amount To Pay at Reception", "", 1, "L", false, 0, "")
	if doctor.Fees == 0 {
		pdf.MultiCell(0, 10, "Please Check at Counter", "", "L", false)
	} else {
		pdf.MultiCell(0, 10, fmt.Sprintf("%v", doctor.Fees), "", "L", false)
	}

	// save pdf to google cloud bucket
	publicUrl, err := puc.savePdfToCloud(ctx, pdf, patient.Name, doctor.Name)
	if err != nil {
		log.Printf("Error %v\n, unable to save pdf to cloud\n\n", err)
	}
	return publicUrl, nil
}