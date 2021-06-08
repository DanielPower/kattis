package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gitlab.com/danielpower/kattis/go/problems"
)

var problemMap = map[string]problems.Problem{
	"growlinggears": problems.GrowlingGears,
	"hello":         problems.Hello,
	"tarifa":        problems.Tarifa,
}

func main() {
	if len(os.Args[1:]) >= 1 {
		problemName := os.Args[1]
		problem := problemMap[problemName]
		runAllTestsForProblem(problem)
	} else {
		for _, problem := range problemMap {
			runAllTestsForProblem(problem)
		}
	}
}

func runAllTestsForProblem(problem problems.Problem) {
	testCountString := strconv.Itoa(len(problem.Tests))
	for testId, fileName := range problem.Tests {
		testIdString := strconv.Itoa(testId + 1)
		fmt.Println("[" + testIdString + "/" + testCountString + "] " + problem.Name)
		runSingleTestForProblem(problem, fileName)
	}
}

func runSingleTestForProblem(problem problems.Problem, fileName string) {
	input, expectedOutput := getTestStrings(fileName)
	output := problem.Run(input)
	if strings.TrimSuffix(output, "\n") == strings.TrimSuffix(expectedOutput, "\n") {
		fmt.Println("✅ Passed")
	} else {
		fmt.Println("❌ Failed")
	}
}

func getTestStrings(testName string) (string, string) {
	var results [2]string
	for i, extenstion := range []string{"in", "ans"} {
		bytes, err := os.ReadFile("tests/" + testName + "." + extenstion)
		if err != nil {
			log.Fatal(err)
		}
		string := strings.TrimSuffix(string(bytes), "\n")
		results[i] = string
	}
	return results[0], results[1]
}
