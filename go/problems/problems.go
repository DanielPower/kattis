package problems

import (
	"io"
)

type Problem struct {
	Name  string
	Run   func(input io.Reader, output io.Writer)
	Tests []string
}
