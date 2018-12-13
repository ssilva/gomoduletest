package gomoduletest

import "fmt"

// Returns a greeting
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
