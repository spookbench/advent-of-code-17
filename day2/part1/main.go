package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	result := 0

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "	")

		min, _ := strconv.Atoi(numbers[0])
		max := 0

		for _, element := range numbers {
			num, _ := strconv.Atoi(element)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		result += (max - min)
	}
	fmt.Println(result)
}
