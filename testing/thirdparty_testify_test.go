package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestify(t *testing.T) {
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
			assert.Equal(t, tt.want, Hello(tt.input))
		})
	}
}
