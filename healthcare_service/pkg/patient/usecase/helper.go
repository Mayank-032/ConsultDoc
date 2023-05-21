package usecase

import (
	"context"
	"fmt"
	"healthcare-service/domain"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func (puc PatientUCase) savePdfToCloud(ctx context.Context, pdf *gofpdf.Fpdf, patientName, doctorName string) (string, error) {
	client := puc.Client
	objName := fmt.Sprintf("%v_%v.pdf", patientName, doctorName)
	bucketName := os.Getenv("ConsultDocGoogleCloudBucket")

	bucket := client.Bucket(bucketName)
	wc := bucket.Object(objName).NewWriter(ctx)
	if err := pdf.Output(wc); err != nil {
		log.Printf("Error %v", err)
		return domain.EmptyString, err
	}
	if err := wc.Close(); err != nil {
		log.Printf("Error %v", err)
		return domain.EmptyString, err
	}
	googleCloudBaseUrl := os.Getenv("GoogleCloudBaseUrl")
	return fmt.Sprintf("%v/%v/%v", googleCloudBaseUrl, bucketName, objName), nil
}