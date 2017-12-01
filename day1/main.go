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

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	numbers := strings.Split(string(dat), "")
	result := 0

	for index, element := range numbers {
		if index+1 >= 0 && index+1 < len(numbers) {
			if element == numbers[index+1] {
				el, _ := strconv.Atoi(element)
				result += el
			}
		} else {
			if element == numbers[0] {
				el, _ := strconv.Atoi(element)
				result += el
			}
		}
	}

	fmt.Print(result)
}
