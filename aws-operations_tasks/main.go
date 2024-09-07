package main

import (
	"aws-Project/userpermission"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	// Initialize a session that loads credentials from ~/.aws/credentials and ~/.aws/config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your region
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}

	// Create an IAM client
	svc := iam.New(sess)

	// Retrieve user permissions
	users, err := userpermission.GetUsers(svc)
	if err != nil {
		log.Fatalf("failed to get users, %v", err)
	}

	// Write users and their permissions to a CSV file
	file, err := os.Create("users_permissions.csv")
	if err != nil {
		log.Fatalf("failed to create file, %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	if err := writer.Write([]string{"UserName", "Permissions"}); err != nil {
		log.Fatalf("failed to write header to CSV, %v", err)
	}

	// Write user details
	for _, user := range users {
		permissions, err := userpermission.GetUserPermissions(svc, *user.UserName)
		if err != nil {
			log.Printf("failed to get permissions for user %s, %v", *user.UserName, err)
			continue
		}
		row := []string{
			*user.UserName,
			permissions,
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("failed to write row to CSV, %v", err)
		}
	}

	fmt.Println("User permissions have been written to users_permissions.csv")
}
