package bootstrap

import (
	"database/sql"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	_patientRepo "healthcare-service/pkg/patient/repository"
	_patientUCase "healthcare-service/pkg/patient/usecase"
	_patientRoutes "healthcare-service/pkg/patient/routes"

	"github.com/gin-gonic/gin"
)

var (
	patientRepo repository.IPatientRepository
	patientUCase usecase.IPatientUseCase
)

func initRepos() {
	patientRepo = _patientRepo.NewPatientRepo(&sql.DB{})
}

func initUCases() {
	patientUCase = _patientUCase.NewPatientUsecase(patientRepo)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_patientRoutes.Init(apiGroup, patientUCase)
}

func Init(apiGroup *gin.RouterGroup) {
	initRepos()
	initUCases()
	initAPIs(apiGroup)
}