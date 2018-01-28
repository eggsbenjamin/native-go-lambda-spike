package main

import "fmt"

type ErrInvalidAge struct {
	Age int
}

func (e ErrInvalidAge) Error() string {
	return fmt.Sprintf("invalid age : %d", e.Age)
}
