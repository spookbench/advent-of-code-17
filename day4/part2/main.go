package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
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
			for i, w := range words {
				if !alreadyFound && index != i {
					w1 := sortString(word)
					w2 := sortString(w)
					fmt.Println(w1, w2)
					if w1 == w2 {
						alreadyFound = true
						invalidLines++
					}
				}
			}
		}
	}
	fmt.Println(invalidLines, numberOfLines)
	result = numberOfLines - invalidLines
	fmt.Println(result)
}
