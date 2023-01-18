package main

import (
	"fmt"

	"github.com/mwelwankuta/adc-challenges/dayone"
)

func main() {
	//  Day 1
	questionOne := dayone.Question{}
	questionOne.LoadInput("input.txt")

	questionOne.CountCalories()
	fmt.Printf("DAY 1, QESTION 1: -> %v \n", questionOne.PartOne)

	questionOne.TopThreeElves()
	fmt.Printf("DAY 1, QESTION 2: -> %v \n", questionOne.PartTwo)

}
