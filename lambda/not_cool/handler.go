package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func NotCool(person *models.Person) error {
	fmt.Printf("Unfortunately %s, due to your hair colour you have been deemed to be not cool by the step function.\n", person.Name)

	return nil
}

func main() {
	lambda.Start(NotCool)
}
