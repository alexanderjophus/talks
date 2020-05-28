package testing

import "fmt"

// Hello greets people
func Hello(name string) (greeting string) {
	return fmt.Sprintf("Hello, %s!", name)
}
