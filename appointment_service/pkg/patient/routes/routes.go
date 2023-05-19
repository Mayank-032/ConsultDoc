package routes

import (
	"healthcare-service/domain/interfaces/usecase"

	"github.com/gin-gonic/gin"
)

var PatientUCase usecase.IPatientUseCase
func Init(apiGroup *gin.RouterGroup, patientUCase usecase.IPatientUseCase) {
	PatientUCase = patientUCase

	apiGroup.POST("/appointment/create", createAppointment)
}

func createAppointment(c *gin.Context) {
}