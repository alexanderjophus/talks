package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open("examples/README.md")
	if err != nil {
		// fmt.Errorf("error opening file: %v", err)
		panic("error opening file")
	}
	defer f.Close()

	if _, err := io.Copy(os.Stdout, f); err != nil {
		// fmt.Errorf("error copying contents: %v", err)
		panic("error copying contents")
	}
}
