package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lineScanner = bufio.NewScanner(os.Stdin)
	lineScanner.Scan()
	testCaseCount, err := strconv.Atoi(lineScanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < testCaseCount; i++ {
		lineScanner.Scan()
		gearCount, err := strconv.Atoi(lineScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		maxTorque := 0
		bestGear := 0
		for j := 0; j < gearCount; j++ {
			gear := [3]int{}
			lineScanner.Scan()
			gearParameters := lineScanner.Text()
			gearScanner := bufio.NewScanner(strings.NewReader(gearParameters))
			gearScanner.Split(bufio.ScanWords)
			for k := 0; k < 3; k++ {
				gearScanner.Scan()
				gearParameter, err := strconv.Atoi(gearScanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				gear[k] = gearParameter
			}
			a, b, c := gear[0], gear[1], gear[2]
			torque := getTorque(a, b, c)
			if torque > maxTorque {
				maxTorque = torque
				bestGear = j
			}
		}
		io.WriteString(os.Stdout, strconv.Itoa(bestGear+1)+"\n")
	}
}

func getTorque(a, b, c int) int {
	r := -b / (2 * -a)
	return -a*(r^2) + (b * r) + c
}
