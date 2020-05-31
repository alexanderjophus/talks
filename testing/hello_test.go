package testing

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "happy path: Alexander", input: "Alexander", want: "Hello, Alexander!"},
		{name: "happy path: empty arg", input: "", want: "Hello, World!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if Hello(test.input) != test.want {
				t.Errorf("want: %s\n got: %s", test.want, Hello(test.name))
			}
		})
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Hello("Alexander")
	}
}

func ExampleHello() {
	greeting := Hello("Alexander")
	fmt.Println(greeting)
	// Output: Hello, Alexander!
}
