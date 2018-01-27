package age_check

import "fmt"

type ErrInvalidAge struct {
	Age int
}

func (e ErrInvalidAge) Error() string {
	return fmt.Sprintf("invalid age : %d", e.Age)
}
