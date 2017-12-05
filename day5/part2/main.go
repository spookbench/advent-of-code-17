package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	check(err)

	return file
}
func main() {
	dat := openFile("../input.txt")
	stepsString := strings.Split(string(dat), "\n")

	currIndex := 0
	nextIndex := 0
	result := 0

	steps := convertToIntArray(stepsString)

	for nextIndex >= 0 && nextIndex < len(steps) {
		result++
		currIndex = nextIndex
		nextIndex += steps[currIndex]

		if steps[currIndex] >= 3 {
			steps[currIndex]--
		} else if steps[currIndex] < 3 {
			steps[currIndex]++
		}
	}
	fmt.Println(result)
}

func convertToIntArray(arr []string) []int {
	result := make([]int, 0, len(arr))
	for _, element := range arr {
		el, _ := strconv.Atoi(element)
		result = append(result, el)
	}
	return result
}
