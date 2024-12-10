package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileBytes
}

func ReadFileLines(path string) []string {
	fileBytes := ReadFileBytes(path)
	lines := []string{}
	separator := []byte("\n")
	for _, line := range bytes.Split(fileBytes, separator) {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}
	return lines
}

func SplitLists(lines []string) [][]int {
	listoflists := [][]int{}
	for _, line := range lines {

		list := []int{}

		linesplit := strings.Split(line, " ")

		for _, i := range linesplit {
			j, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, j)
		}

		listoflists = append(listoflists, list)
	}
	return listoflists
}

func CheckListSafety(list []int) int {
	i := 0
	direction := 0
	if list[0] < list[1] {
		direction = 1
	} else if list[0] > list[1] {
		direction = -1
	}
	if direction != 0 {
		for i < len(list)-1 {
			if direction == 1 {
				if list[i+1] > list[i] && list[i+1] < list[i]+4 {
					i = i + 1
				} else {
					i = len(list) + 1
				}
			} else {
				if list[i+1] < list[i] && list[i+1] > list[i]-4 {
					i = i + 1
				} else {
					i = len(list) + 1
				}
			}
		}
		if i != len(list)+1 {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func remove(slice []int, i int) []int {
	var result []int
	for id, element := range slice {
		if id != i {
			result = append(result, element)
		}
	}
	return result
}

func PartTwo(lines []string) int {
	listoflists := SplitLists(lines)
	var count int
	for _, list := range listoflists {
		if CheckListSafety(list) == 1 {
			count = count + 1
		} else {
			i := 0
			for i < len(list) {
				if CheckListSafety(remove(list, i)) == 1 {
					count = count + 1
					i = len(list)
				} else {
					i = i + 1
				}
			}
		}
	}
	return count
}

func PartOne(lines []string) int {
	listoflists := SplitLists(lines)
	var count int
	for _, list := range listoflists {
		count = count + CheckListSafety(list)
	}
	return count
}

func main() {
	lines := ReadFileLines("../../inputs/day02/part1.txt")
	partOne := PartOne(lines)
	fmt.Println(partOne)
	partTwo := PartTwo(lines)
	fmt.Println(partTwo)
}
