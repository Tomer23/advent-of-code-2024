package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func SplitLists(lines []string) ([]int, []int) {
	listOne := []int{}
	listTwo := []int{}

	for _, line := range lines {
		// |3   4
		// ["3","4"]
		// 3
		// 4
		numbers := strings.Split(line, "   ")
		numOne, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		numTwo, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		listOne = append(listOne, numOne)
		listTwo = append(listTwo, numTwo)
	}
	return listOne, listTwo
}

func SplitListsAndMap(lines []string) ([]int, map[int]int) {
	listOne := []int{}
	listTwoCounts := make(map[int]int)

	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		numOne, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		numTwo, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		listTwoCounts[numTwo] += 1
		listOne = append(listOne, numOne)
	}
	return listOne, listTwoCounts
}

func CreateNumberMap(listOne []int, mapTwo map[int]int) int {
	similarityScore := 0
	for _, numOne := range listOne {
		score := numOne * mapTwo[numOne]
		similarityScore += score
	}
	return similarityScore
}

func PartOne(lines []string) int {
	listOne, listTwo := SplitLists(lines)
	sort.Ints(listOne)
	sort.Ints(listTwo)
	totalScore := 0
	for i := range listOne {
		totalScore += int(math.Abs(float64(listOne[i] - listTwo[i])))
	}
	return totalScore
}

func PartTwo(lines []string) int {
	listOne, mapTwo := SplitListsAndMap(lines)
	// fmt.Println(listOne, mapTwo)
	similarityScore := CreateNumberMap(listOne, mapTwo)
	return similarityScore
}

func main() {
	lines := ReadFileLines("../../inputs/day01/part2.txt")
	partOne := PartOne(lines)
	fmt.Println(partOne)
	partTwo := PartTwo(lines)
	fmt.Println(partTwo)
}
