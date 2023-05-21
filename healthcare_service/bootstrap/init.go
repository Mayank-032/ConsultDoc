package bootstrap

import (
	DB "healthcare-service/db"
	"healthcare-service/domain/interfaces"
	"healthcare-service/domain/interfaces/controller"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	_commonRepo "healthcare-service/pkg/common/repository"
	_commonRoutes "healthcare-service/pkg/common/routes"
	_commonUCase "healthcare-service/pkg/common/usecase"
	_patientController "healthcare-service/pkg/patient/controller"
	_patientRepo "healthcare-service/pkg/patient/repository"
	_patientUCase "healthcare-service/pkg/patient/usecase"
	"healthcare-service/rabbitmq"
	_consumer "healthcare-service/rabbitmq/consumer"

	"github.com/gin-gonic/gin"
)

var (
	commonRepo        repository.ICommonRepository
	commonUCase       usecase.ICommonUseCase
	patientRepo       repository.IPatientRepository
	patientUCase      usecase.IPatientUseCase
	patientController controller.IPatientController
	consumer          interfaces.IConsumer
)

func initRepos() {
	commonRepo = _commonRepo.NewCommonRepo(DB.Client)
	patientRepo = _patientRepo.NewPatientRepo(DB.Client)
}

func initUCases() {
	commonUCase = _commonUCase.NewCommonUCase(commonRepo)
	patientUCase = _patientUCase.NewPatientUCase(DB.CloudClient)
}

func initControllers() {
	patientController = _patientController.NewPatientController(rabbitmq.Conn, patientUCase, patientRepo)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_commonRoutes.Init(apiGroup, commonUCase)
}

func initConsumer() {
	consumer = _consumer.NewConsumerLayer(patientController)
}

func Init(apiGroup *gin.RouterGroup) {
	initRepos()
	initControllers()
	initUCases()
	initAPIs(apiGroup)

	initConsumer()
	consumer.StartConsumers()
}
