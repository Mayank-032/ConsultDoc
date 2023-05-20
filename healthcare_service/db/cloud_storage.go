package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var CloudClient *storage.Client
func ConnectCloud(credentialFilePath string) {
	ctx := context.Background()

    cloudClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
    if err != nil {
        log.Printf("failed to create client: %v\n\n", err)
		return
    }
	fmt.Println("Connected to Google Cloud Store!")
	CloudClient = cloudClient
}