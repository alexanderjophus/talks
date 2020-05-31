package testing

import (
	"fmt"
)

// Hello greets people
func Hello(name string) (greeting string) {
	// time.Sleep(20 * time.Millisecond)
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
