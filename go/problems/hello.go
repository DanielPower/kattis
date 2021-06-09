package problems

import (
	"io"
)

var Hello = Problem{
	Name:  "hello",
	Run:   runHello,
	Tests: []string{"hello"},
}

func runHello(input io.Reader, output io.Writer) {
	io.WriteString(output, "Hello World!")
}
