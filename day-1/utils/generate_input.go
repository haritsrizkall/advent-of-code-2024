package utils

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GenerateInput() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	leftPair := []int{}
	rightPair := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		lineElements := strings.Split(line, "   ")
		ele1, err := strconv.Atoi(lineElements[0])
		if err != nil {
			panic(err)
		}
		ele2, err := strconv.Atoi(lineElements[1])
		if err != nil {
			panic(err)
		}
		leftPair = append(leftPair, ele1)
		rightPair = append(rightPair, ele2)
	}
	sort.Ints(leftPair)
	sort.Ints(rightPair)

	return leftPair, rightPair
}
