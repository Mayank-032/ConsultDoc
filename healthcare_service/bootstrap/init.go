package bootstrap

import (
	DB "healthcare-service/db"
	"healthcare-service/domain"
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
	_doctorController "healthcare-service/pkg/doctor/controller"
	_doctorUCase "healthcare-service/pkg/doctor/usecase"
	_doctorRepo "healthcare-service/pkg/doctor/repository"
	"healthcare-service/rabbitmq"
	_consumer "healthcare-service/rabbitmq/consumer"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	commonRepo        repository.ICommonRepository
	commonUCase       usecase.ICommonUseCase
	patientRepo       repository.IPatientRepository
	patientUCase      usecase.IPatientUseCase
	patientController controller.IPatientController
	doctorRepo repository.IDoctorRepository
	doctorUCase usecase.IDoctorUseCase
	doctorController controller.IDoctorController
	consumer          interfaces.IConsumer
	smtpCreds         domain.SMTP_Cred
)

func initRepos() {
	commonRepo = _commonRepo.NewCommonRepo(DB.Client)
	patientRepo = _patientRepo.NewPatientRepo(DB.Client)
	doctorRepo = _doctorRepo.NewDoctorRepo(DB.Client)
}

func initUCases(smtpCred domain.SMTP_Cred) {
	commonUCase = _commonUCase.NewCommonUCase(commonRepo)
	patientUCase = _patientUCase.NewPatientUCase(DB.CloudClient)
	doctorUCase = _doctorUCase.NewDoctorUCase(smtpCred, doctorRepo)
}

func initControllers() {
	patientController = _patientController.NewPatientController(rabbitmq.Conn, patientUCase, patientRepo)
	doctorController = _doctorController.NewDoctorController(doctorUCase)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_commonRoutes.Init(apiGroup, commonUCase)
}

func initConsumer() {
	consumer = _consumer.NewConsumerLayer(patientController)
}

func Init(apiGroup *gin.RouterGroup) {
	smtpCreds = loadSMTPCredentials()

	initRepos()
	initControllers()
	initUCases(smtpCreds)
	initAPIs(apiGroup)

	initConsumer()
	consumer.StartConsumers()
}

func loadSMTPCredentials() domain.SMTP_Cred {
	return domain.SMTP_Cred{
		BusinessEmail: os.Getenv("BusinessEmail"),
		BusinessPassword: os.Getenv("BusinessPassword"),
		SMTPHost: os.Getenv("SMTPHost"),
		SMTPPort: os.Getenv("SMTPPort"),
	}
}