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
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)
	numbers := strings.Split(string(dat), "")
	interval := len(numbers) / 2
	firstHalf := numbers[interval:]
	result := 0

	for index, _ := range firstHalf {
		if numbers[index] == numbers[index+interval] {
			el, _ := strconv.Atoi(numbers[index])
			result += el
		}
	}
	fmt.Print(result * 2)
}
