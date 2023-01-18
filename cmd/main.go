package main

import (
	"fmt"

	"github.com/mwelwankuta/advent-of-code-go-2022/internal/dayone"
	"github.com/mwelwankuta/advent-of-code-go-2022/internal/daytwo"
)

func main() {
	//  Day 1
	dayOne := dayone.Question{}
	dayOne.LoadInput("dayone", "input.txt")

	dayOne.CountCalories()
	dayOne.TopThreeElves()

	fmt.Printf("DAY 1, QESTION 1: -> %v \n", dayOne.PartOne)
	fmt.Printf("DAY 1, QESTION 2: -> %v \n\n", dayOne.PartTwo)

	// Day 2
	dayTwo := daytwo.Question{}
	dayTwo.LoadInput("daytwo", "input.txt")
	dayTwo.CalculateRPSTotal()

	// part 2
	dayTwo.LoadInput("daytwo", "input.two.txt")
	dayTwo.CalculateRoundOutcomes()

	fmt.Printf("DAY 2, QESTION 1: -> %v \n", dayTwo.PartOne)
	fmt.Printf("DAY 2, QESTION 2: -> %v \n\n", dayTwo.PartTwo)

}
