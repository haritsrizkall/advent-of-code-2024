package main

import (
	"firstday/parts"
	"firstday/utils"
	"fmt"
)

func main() {
	leftPair, rightPair := utils.GenerateInput()

	res := parts.PartTwo(leftPair, rightPair)

	fmt.Printf("Result = %d", res)
}
