package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func HasGingerFriends(person *models.Person) (*models.Person, error) {
	for _, friend := range person.Friends {
		if friend.HairColour == "ginger" {
			return person, nil
		}
	}

	return person, ErrNoGingerFriends{person.Name}
}

func main() {
	lambda.Start(HasGingerFriends)
}
