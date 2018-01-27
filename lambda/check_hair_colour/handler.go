package age_check

import "github.com/eggsbenjamin/native-go-lambda-spike/models"

func HairColourCheck(person *models.Person) (*models.Person, error) {
	if person.HairColour != "ginger" {
		return person, ErrHairColour{person.HairColour}
	}

	return person, nil
}
