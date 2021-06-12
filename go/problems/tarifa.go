package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var lineScanner = bufio.NewScanner(os.Stdin)
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
	io.WriteString(os.Stdout, strconv.Itoa(totalData+monthlyData))
}

func panicIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
