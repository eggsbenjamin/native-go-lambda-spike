package age_check

import (
	"fmt"

	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func UnderAge(person *models.Person) error {
	fmt.Printf("Unfortunately %s, you have been deemed to be too young by the step function.\n")

	return nil
}
