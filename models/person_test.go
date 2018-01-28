package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	tests := []struct {
		title   string
		rawJSON []byte
		object  *Person
	}{
		{
			"non-existant friends",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown" }`),
			&Person{"Bob", 22, "brown", []*Friend{}},
		},
		{
			"null friends",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : null }`),
			&Person{"Bob", 22, "brown", []*Friend{}},
		},
		{
			"zero friends",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : [] }`),
			&Person{"Bob", 22, "brown", []*Friend{}},
		},
		{
			"one friend",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : [31,139,8,0,0,0,0,0,0,255,138,230,82,80,168,86,80,202,75,204,77,85,82,176,82,80,10,78,204,201,169,84,210,81,80,74,76,7,11,152,152,232,40,40,101,36,102,22,197,39,231,231,228,151,22,129,21,165,103,230,165,167,22,129,84,165,21,101,166,230,165,20,131,68,163,99,107,185,98,185,0,1,0,0,255,255,94,1,75,165,80,0,0,0]}`),
			&Person{"Bob", 22, "brown", []*Friend{
				&Friend{"Sally", 44, "ginger", []*Friend{}}},
			},
		},
		{
			"multiple friends",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : [31,139,8,0,0,0,0,0,0,255,138,230,82,80,168,86,80,202,75,204,77,85,82,176,82,80,114,202,79,82,210,81,80,74,76,7,115,141,140,116,20,148,50,18,51,139,226,147,243,115,242,75,139,192,74,146,138,242,203,243,64,138,210,138,50,83,243,82,138,65,130,209,177,181,58,10,232,70,101,230,228,32,153,101,108,140,205,172,156,252,188,148,84,34,12,11,78,204,201,169,68,50,205,196,4,139,105,233,153,121,233,169,69,152,166,113,197,114,1,2,0,0,255,255,89,194,121,8,232,0,0,0] }`),
			&Person{"Bob", 22, "brown", []*Friend{
				&Friend{"Bob", 22, "brown", []*Friend{}},
				&Friend{"Bill", 33, "blonde", []*Friend{}},
				&Friend{"Sally", 44, "ginger", []*Friend{}},
			}},
		},
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			var p *Person
			require.NoError(t, json.Unmarshal(v.rawJSON, &p))
			require.Equal(t, v.object, p)

			output, err := json.Marshal(p)
			require.NoError(t, err)

			var p2 *Person
			require.NoError(t, json.Unmarshal(output, &p2))
			require.Equal(t, v.object, p2)
		})
	}
}
