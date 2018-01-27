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
			title: "< 18",
			input: &models.Person{
				Age: 17,
			},
			expectedOutput: &models.Person{
				Age: 17,
			},
			expectedError: ErrInvalidAge{17},
		},
		{
			title: "18",
			input: &models.Person{
				Age: 18,
			},
			expectedOutput: &models.Person{
				Age: 18,
			},
			expectedError: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			output, err := AgeCheck(v.input)

			require.Equal(t, v.expectedOutput, output)
			require.Equal(t, v.expectedError, err)
		})
	}
}
