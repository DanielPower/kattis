package utils

import (
	"log"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PanicIfErr(err error) {
	if err != nil {
		log.Fatal()
	}
}
