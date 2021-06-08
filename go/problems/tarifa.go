package problems

import (
	"bufio"
	"strconv"
	"strings"

	"gitlab.com/danielpower/kattis/go/utils"
)

var Tarifa = Problem{
	Name: "tarifa",
	Run: runTarifa,
	Tests: []string{"tarifa.1", "tarifa.2", "tarifa.3"},
}

func runTarifa(input string) string {
	var lineScanner = bufio.NewScanner(strings.NewReader(input))
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
	return strconv.Itoa(totalData + monthlyData)
}
