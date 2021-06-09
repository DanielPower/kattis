package problems

import (
	"bufio"
	"io"
	"strconv"

	"gitlab.com/danielpower/kattis/go/utils"
)

var Tarifa = Problem{
	Name:  "tarifa",
	Run:   runTarifa,
	Tests: []string{"tarifa.1", "tarifa.2", "tarifa.3"},
}

func runTarifa(input io.Reader, output io.Writer) {
	var lineScanner = bufio.NewScanner(input)
	lineScanner.Scan()
	monthlyDataString := lineScanner.Text()
	monthlyData, err := strconv.Atoi(monthlyDataString)
	utils.PanicIfErr(err)
	totalData := 0
	lineScanner.Scan()
	for lineScanner.Scan() {
		usageString := lineScanner.Text()
		usage, err := strconv.Atoi(usageString)
		utils.PanicIfErr(err)
		totalData = totalData + monthlyData - usage
	}
	io.WriteString(output, strconv.Itoa(totalData+monthlyData))
}
