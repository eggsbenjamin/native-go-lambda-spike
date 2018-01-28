package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func UnderAge(person *models.Person) error {
	fmt.Printf("Unfortunately %s, you have been deemed to be too young by the step function.\n", person.Name)

	return nil
}

func main() {
	lambda.Start(UnderAge)
}
