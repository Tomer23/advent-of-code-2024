package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileString(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileBytes)
}

func ProcessOrders(input string, order bool) bool {
	var dos []int
	var donts []int
	start, i := 0, 0
	for i >= 0 {
		i = strings.Index(input[start:], "do()")
		if i > 0 {
			dos = append(dos, i)
			start = i + 1
		}
	}

	start, i = 0, 0
	for i >= 0 {
		i = strings.Index(input[start:], "don't()")
		if i > 0 {
			donts = append(donts, i)
			start = i + 1
		}
	}

	if len(dos) == 0 && len(donts) == 0 {
		return order
	} else if len(dos) == 0 {
		return false
	} else if len(donts) == 0 {
		return true
	} else {
		if dos[len(dos)-1] > donts[len(donts)-1] {
			return true
		} else {
			return false
		}
	}

}

func PartOne(input string) int {
	count := 0
	s := strings.SplitAfter(input, "mul(")
	for _, element := range s {
		first_pair := strings.Split(element, ")")[0]
		num_pair := strings.Split(first_pair, ",")
		if len(num_pair) > 1 {
			i, err1 := strconv.Atoi(num_pair[0])
			j, err2 := strconv.Atoi(num_pair[1])

			if err1 == nil && err2 == nil {
				count = count + i*j
			}
		}
	}
	return count
}

func PartTwo(input string) int {
	count := 0
	s := strings.SplitAfter(input, "mul(")
	order := ProcessOrders(s[0], true)
	for _, element := range s {
		first_pair := strings.Split(element, ")")[0]
		num_pair := strings.Split(first_pair, ",")
		if len(num_pair) > 1 {
			i, err1 := strconv.Atoi(num_pair[0])
			j, err2 := strconv.Atoi(num_pair[1])

			if err1 == nil && err2 == nil {
				if order {
					count = count + i*j
				}
			}
		}
		order = ProcessOrders(element, order)
	}
	return count
}

func main() {
	input := ReadFileString("../../inputs/day03/part1.txt")
	partOne := PartOne(input)
	fmt.Println(partOne)
	partTwo := PartTwo(input)
	fmt.Println(partTwo)
}
