package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func Cool(person *models.Person) error {
	fmt.Printf("Congratulations %s! You have been deemed to be cool by the step function!\n", person.Name)

	return nil
}

func main() {
	lambda.Start(Cool)
}
