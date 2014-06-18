package infrastructure

import (
	"fmt"
)

type Logger struct{}

// implements usecase.Logger
// not used in usecase_test.go as the usecase package may
// not depend on this package infrastructure.
func (logger Logger) Log(message string) error {
	fmt.Println("Log message: " + message)
	return nil
}
