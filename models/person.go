package models

type Person struct {
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	HairColour string    `json:"hair_colour"`
	Friends    []*Friend `json:"friends"`
}

type Friend Person
