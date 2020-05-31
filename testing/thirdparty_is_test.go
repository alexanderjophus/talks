package testing

import (
	"testing"

	"github.com/matryer/is"
)

func TestIs(t *testing.T) {
	tts := []struct {
		name  string
		input string
		want  string
	}{
		{name: "happy path: Alexander", input: "Alexander", want: "Hello, Alexander!"},
		{name: "happy path: empty arg", input: "", want: "Hello, World!"},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			greeting := Hello(tt.input)
			is.Equal(greeting, tt.want) // contextual comment
		})
	}
}
