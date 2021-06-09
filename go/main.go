package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gitlab.com/danielpower/kattis/go/problems"
	"gitlab.com/danielpower/kattis/go/utils"
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

func runSingleTestForProblem(problem problems.Problem, testName string) {
	inputFile, err := os.Open("tests/" + testName + ".in")
	utils.PanicIfErr(err)
	expectedOutput, err := ioutil.ReadFile("tests/" + testName + ".ans")
	utils.PanicIfErr(err)
	var outputBuffer bytes.Buffer
	outputWriter := bufio.NewWriter(&outputBuffer)
	problem.Run(inputFile, outputWriter)
	outputWriter.Flush()
	if strings.TrimSuffix(outputBuffer.String(), "\n") == strings.TrimSuffix(string(expectedOutput), "\n") {
		fmt.Println("✅ Passed")
	} else {
		fmt.Println("❌ Failed")
	}
}
