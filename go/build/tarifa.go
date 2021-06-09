package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runTarifa(os.Stdin, os.Stdout)
}

func runTarifa(input io.Reader, output io.Writer) {
	var lineScanner = bufio.NewScanner(input)
	lineScanner.Scan()
	monthlyDataString := lineScanner.Text()
	monthlyData, err := strconv.Atoi(monthlyDataString)
	panicIfErr(err)
	totalData := 0
	lineScanner.Scan()
	for lineScanner.Scan() {
		usageString := lineScanner.Text()
		usage, err := strconv.Atoi(usageString)
		panicIfErr(err)
		totalData = totalData + monthlyData - usage
	}
	io.WriteString(output, strconv.Itoa(totalData+monthlyData))
}
