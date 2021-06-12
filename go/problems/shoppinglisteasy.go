package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Scan() // Ignore first line
	lineScanner.Scan()
	initialItems := getItemsFromLine(lineScanner.Text())
	for lineScanner.Scan() {
		newItems := getItemsFromLine(lineScanner.Text())
		for item := range initialItems {
			if _, exists := newItems[item]; !exists {
				delete(initialItems, item)
			}
		}
	}
	itemCount := strconv.Itoa(len(initialItems))
	fmt.Println(itemCount)
	sortedItems := getSortedItemsFromMap(initialItems)
	for _, item := range sortedItems {
		fmt.Println(item)
	}
}

func getItemsFromLine(line string) map[string]bool {
	var items = make(map[string]bool)
	wordScanner := bufio.NewScanner(strings.NewReader(line))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		items[wordScanner.Text()] = true
	}
	return items
}

func getSortedItemsFromMap(items map[string]bool) []string {
	keys := make([]string, 0)
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
