package age_check

import (
	"fmt"

	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func UnderAge(person *models.Person) error {
	fmt.Printf("Unfortunately %s, due to your hair colour you have been deemed to be not cool by the step function.\n")

	return nil
}
