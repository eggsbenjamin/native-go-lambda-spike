package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func AgeCheck(person *models.Person) (*models.Person, error) {
	if person.Age < 18 {
		return person, ErrInvalidAge{person.Age}
	}

	return person, nil
}

func main() {
	lambda.Start(AgeCheck)
}
