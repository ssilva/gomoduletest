package gomoduletest

import "fmt"

// Returns a greeting
func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
