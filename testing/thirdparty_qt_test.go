package testing

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestQt(t *testing.T) {
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
			c := qt.New(t)
			greeting := Hello(tt.input)
			c.Assert(greeting, qt.Equals, tt.want)
		})
	}
}
