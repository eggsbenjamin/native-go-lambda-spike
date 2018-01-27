package age_check

import (
	"testing"

	"github.com/eggsbenjamin/native-go-lambda-spike/models"
	"github.com/stretchr/testify/require"
)

func TestAgeCheck(t *testing.T) {
	tests := []struct {
		title          string
		input          *models.Person
		expectedOutput *models.Person
		expectedError  error
	}{
		{
			title: "not ginger",
			input: &models.Person{
				HairColour: "blonde",
			},
			expectedOutput: &models.Person{
				HairColour: "blonde",
			},
			expectedError: ErrHairColour{"blonde"},
		},
		{
			title: "ginger",
			input: &models.Person{
				HairColour: "ginger",
			},
			expectedOutput: &models.Person{
				HairColour: "ginger",
			},
			expectedError: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			output, err := HairColourCheck(v.input)

			require.Equal(t, v.expectedOutput, output)
			require.Equal(t, v.expectedError, err)
		})
	}
}
