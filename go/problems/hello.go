package problems

var Hello = Problem{
	Name:  "hello",
	Run:   runHello,
	Tests: []string{"hello"},
}

func runHello(input string) string {
	return "Hello World!"
}
