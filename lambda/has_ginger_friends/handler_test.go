package main

import (
	"testing"

	"github.com/eggsbenjamin/native-go-lambda-spike/models"
	"github.com/stretchr/testify/require"
)

func TestHasGingerFriends(t *testing.T) {
	tests := []struct {
		title          string
		input          *models.Person
		expectedOutput *models.Person
		expectedError  error
	}{
		{
			"no ginger friends",
			&models.Person{"Bob", 33, "brown", []*models.Friend{&models.Friend{"Bill", 34, "blonde", nil}}},
			&models.Person{"Bob", 33, "brown", []*models.Friend{&models.Friend{"Bill", 34, "blonde", nil}}},
			ErrNoGingerFriends{"Bob"},
		},
		{
			"has ginger friends",
			&models.Person{"Bob", 33, "brown", []*models.Friend{&models.Friend{"Bill", 34, "ginger", nil}}},
			&models.Person{"Bob", 33, "brown", []*models.Friend{&models.Friend{"Bill", 34, "ginger", nil}}},
			nil,
		},
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			output, err := HasGingerFriends(v.input)

			require.Equal(t, v.expectedOutput, output)
			require.Equal(t, v.expectedError, err)
		})
	}
}
