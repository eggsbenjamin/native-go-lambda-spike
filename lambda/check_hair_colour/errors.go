package main

type ErrHairColour struct {
	HairColour string
}

func (e ErrHairColour) Error() string {
	return "invalid hair colour : " + e.HairColour
}
