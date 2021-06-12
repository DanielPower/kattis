package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var problems map[string][]string
	data, err := os.ReadFile("problems.json")
	panicIfErr(err)
	err = json.Unmarshal(data, &problems)
	panicIfErr(err)
	for problem, tests := range(problems) {
		runAllTestsForProblem(problem, tests)
	}
}

func runAllTestsForProblem(problem string, tests []string) {
	testCountString := strconv.Itoa(len(tests))
	for testId, fileName := range tests {
		testIdString := strconv.Itoa(testId + 1)
		fmt.Println("[" + testIdString + "/" + testCountString + "] " + problem)
		runSingleTestForProblem(problem, fileName)
	}
}

func runSingleTestForProblem(problem string, fileName string) {
	inputFile, err := os.Open("tests/" + fileName + ".in")
	var output bytes.Buffer
	outputWriter := io.Writer(&output)
	panicIfErr(err)
	expectedOutput, err := os.ReadFile("tests/" + fileName + ".ans")
	panicIfErr(err)
	cmd := exec.Command("go", "run", "problems/"+problem+".go")
	cmd.Stdin = inputFile
	cmd.Stdout = outputWriter
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	panicIfErr(err)
	if strings.TrimSuffix(output.String(), "\n") == strings.TrimSuffix(string(expectedOutput), "\n") {
		fmt.Println("✅ Passed")
	} else {
		fmt.Println("❌ Failed")
	}
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
