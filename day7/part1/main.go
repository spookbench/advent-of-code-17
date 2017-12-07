package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	check(err)

	return file
}

func main() {
	dat := openFile("../input.txt")
	scanner := bufio.NewScanner(dat)

	// result := 0
	var childrenArr []string
	var parentsArr []string

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		parentsArr = append(parentsArr, input[0])

		if len(input) > 2 {
			var children string
			var orphanage []string

			for i := 3; i < len(input); i++ {
				children += input[i]
			}
			orphanage = strings.Split(children, ",")
			for _, e := range orphanage {
				childrenArr = append(childrenArr, e)
			}
		}
	}
	for _, e := range parentsArr {
		if !contains(childrenArr, e) {
			fmt.Println("root element is: ", e)
		}
	}
}

func contains(arr []string, e string) bool {
	for _, element := range arr {
		if element == e {
			return true
		}
	}
	return false
}
