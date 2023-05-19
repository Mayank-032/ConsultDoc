package main

import (
	"healthcare-service/db"
	"healthcare-service/domain"
	"healthcare-service/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error: %v, unable to load .env file", err)
		return
	}
	
	dbCred := getCredentials()
	err = db.Init(dbCred)
	if err != nil {
		log.Printf("Error: %v, unable to connect to database", err.Error())
		return
	}

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":8000")
}

func getCredentials() domain.DBCred {
	return domain.DBCred{
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		Hostname: os.Getenv("hostname"),
		DBName:   os.Getenv("dbname"),
	}
}