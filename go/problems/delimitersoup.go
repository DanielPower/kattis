package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	expectedCloser := []rune{}
	for index, char := range scanner.Text() {
		switch char {
		case '[':
			expectedCloser = append(expectedCloser, ']')
		case '(':
			expectedCloser = append(expectedCloser, ')')
		case '{':
			expectedCloser = append(expectedCloser, '}')
		case ']', ')', '}':
			{
				var closer rune
				var err error
				closer, expectedCloser, err = pop(expectedCloser)
				if err != nil || char != closer {
					fmt.Println(string(char) + " " + strconv.Itoa(index))
					os.Exit(0)
				}
			}
		}
	}
	fmt.Println("ok so far")
}

func pop(arr []rune) (rune, []rune, error) {
	if len(arr) > 0 {
		return arr[len(arr)-1], arr[:len(arr)-1], nil
	}
	return 0, nil, errors.New("EMPTY")
}
