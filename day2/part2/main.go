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

func checkDivision(a int, b int) (bool, int) {
	result := 0
	isDivisible := false

	if a%b == 0 {
		result = a / b
		isDivisible = true
	}
	if b%a == 0 {
		result = b / a
		isDivisible = true
	}

	return isDivisible, result
}

func main() {
	dat := openFile("../input.txt")
	scanner := bufio.NewScanner(dat)

	result := 0

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "	")

		for index, element := range numbers {
			num1, _ := strconv.Atoi(element)
			for i, el := range numbers {
				if index != i {
					num2, _ := strconv.Atoi(el)
					isDivisible, r := checkDivision(num1, num2)
					if isDivisible {
						result += r
					}
				}
			}
		}
	}
	fmt.Println(result / 2)
}
