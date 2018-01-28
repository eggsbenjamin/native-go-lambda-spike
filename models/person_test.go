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
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : "BCJNGGRwuVAAAIBbCiAgeyAibmFtZSIgOiAiU2FsbHkiLCAiYWdlIiA6IDQ0LCAiaGFpcl9jb2xvdXIiIDogImdpbmdlciIsICJmcmllbmRzIiA6IFtdfQpdCgAAAAAzXXbM" }`),
			&Person{"Bob", 22, "brown", []*Friend{
				&Friend{"Sally", 44, "ginger", []*Friend{}}},
			},
		},
		{
			"multiple friends",
			[]byte(`{ "name" : "Bob", "age" : 22, "hair_colour" : "brown", "friends" : "BCJNGGRwuZoAAADxClsKICB7ICJuYW1lIiA6ICJCb2IiLCAiYWcPAPEBMjIsICJoYWlyX2NvbG91ciMAUGJyb3duJQBwZnJpZW5kcxUAXFtdfSwgSwBGaWxsIkwATjMzLCBMAFBsb25kZSYAT2ZyaWVNAAhQU2FsbHkoAJ9hZ2UiIDogNDROAABgZ2luZ2VyJgDwAmZyaWVuZHMiIDogW119Cl0KAAAAACIGjWM=" }`),
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
