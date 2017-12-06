package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
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
	arr := strings.Split(string(dat), "	")

	var prevStates [][]int
	result := 0

	banks := convertToIntArray(arr)
	for !stateAlreadySeen(prevStates, banks) {
		result++
		prevStates = append(prevStates, append([]int(nil), banks...))
		max, maxIndex := maxValue(banks)
		l := len(banks)
		toDistribute := max / l
		rest := max - (toDistribute * l)
		banks[maxIndex] = 0

		if toDistribute == 0 {
			pos := maxIndex + 1
			for i := 0; i < max; i++ {
				if pos >= 0 && pos < len(banks) {
					banks[pos]++
					pos++
				} else {
					pos = 0
					banks[pos]++
					pos++
				}
			}
		} else {
			banks[maxIndex] = rest
			for i, _ := range banks {
				if i != maxIndex {
					banks[i] += toDistribute
				}
			}
		}
	}
	fmt.Println(result)
}

func stateAlreadySeen(a [][]int, b []int) bool {
	result := false
	for _, state := range a {
		if reflect.DeepEqual(state, b) {
			result = true
			break
		}
	}
	return result
}

func convertToIntArray(arr []string) []int {
	result := make([]int, 0, len(arr))
	for _, element := range arr {
		el, _ := strconv.Atoi(element)
		result = append(result, el)
	}
	return result
}

func maxValue(arr []int) (int, int) {
	value := 0
	index := 0
	sortedArr := append([]int(nil), arr...)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedArr)))
	value = sortedArr[0]
	for i, e := range arr {
		if value == e {
			index = i
			break
		}
	}
	return value, index
}
