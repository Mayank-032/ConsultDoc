package routes

import (
	"healthcare-service/domain/interfaces/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CommonUCase usecase.ICommonUseCase
func Init(apiGroup *gin.RouterGroup, commonUCase usecase.ICommonUseCase) {
	CommonUCase = commonUCase

	/*
		-> This API will return an assure patient that appointment creation is in progress and
		in sometime they will get their receipt

		-> Since we are processing request in rabbitMQ
	*/
	apiGroup.GET("/doctor/list", fetchDoctorsList)
}

func fetchDoctorsList(c *gin.Context) {
	resData := gin.H{"status": false}
	request := FetchDoctorListRequest{}
	
	err := c.ShouldBindJSON(&request)
	if err == nil {
		err = request.Validate()
	}
	if err != nil {
		log.Printf("Error: %v\n, invalid_request\n\n", err.Error())
		resData["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, resData)
		return
	}

	address := request.toAddressDto()
	doctors, err := CommonUCase.FetchDoctorsList(c, address)
	if err != nil {
		log.Printf("Error: %v\n. unable_to_doctor's_list\n\n", err.Error())
		resData["message"] = "unable to fetch doctor's list"
		c.JSON(http.StatusOK, resData)
		return
	}
	resData["status"] = true
	resData["message"] = "successfully fetched list"
	resData["doctors"] = doctors
	c.JSON(http.StatusOK, resData)
}