package bootstrap

import (
	DB "healthcare-service/db"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	_patientRepo "healthcare-service/pkg/common/repository"
	_patientRoutes "healthcare-service/pkg/common/routes"
	_patientUCase "healthcare-service/pkg/common/usecase"

	"github.com/gin-gonic/gin"
)

var (
	commonRepo  repository.ICommonRepository
	commonUCase usecase.ICommonUseCase
)

func initRepos() {
	commonRepo = _patientRepo.NewCommonRepo(DB.Client)
}

func initUCases() {
	commonUCase = _patientUCase.NewCommonUCase(commonRepo)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_patientRoutes.Init(apiGroup, commonUCase)
}

func Init(apiGroup *gin.RouterGroup) {
	initRepos()
	initUCases()
	initAPIs(apiGroup)
}
