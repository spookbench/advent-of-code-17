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

	numberOfLines := 0
	invalidLines := 0
	result := 0

	for scanner.Scan() {
		numberOfLines++
		passphrase := scanner.Text()
		words := strings.Split(passphrase, " ")
		alreadyFound := false

		for index, word := range words {
			if alreadyFound {
				break
			}
			for i, w := range words {
				if index != i && w == word {
					if !alreadyFound {
						invalidLines++
						alreadyFound = true
					}
				}
			}
		}
	}
	result = numberOfLines - invalidLines
	fmt.Println(result)
}
