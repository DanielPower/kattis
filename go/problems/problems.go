package problems

type Problem struct {
	Name  string
	Run   func(input string) string
	Tests []string
}
