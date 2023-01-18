// Package dayone provides solution to the first day of the advent of code
// challenges.
package dayone

import (
	"sort"
	"strconv"
	"strings"

	"github.com/mwelwankuta/advent-of-code-go-2022/internal/question"
)

type Question struct {
	question.Question
}

// max(a, b string)
// returns largest value between the passed
// e.g LoadInput("input.txt")
// the function searches the current directory for a file and panics when non is found
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// CountCalories()
// the function calculates the elf with the food that contains the most calories
func (q *Question) CountCalories() {
	largestCalorieCount := 0

	elfFood := strings.Split(q.File, "\n\n")
	for _, food := range elfFood {
		currentFood := strings.Split(food, "\n")

		currentLargest := 0
		for _, calorie := range currentFood {
			if calorie == "" {
				continue
			}

			currentCalorie, err := strconv.Atoi(calorie)
			if err != nil {
				panic(err)
			}
			currentLargest += currentCalorie

			largestCalorieCount = max(largestCalorieCount, currentLargest)
		}
	}

	q.PartOne = largestCalorieCount
}

// sumOfSlice(slice []int)
// returns the sum of elements
// e.g sumOfSlice([]{1,2,3,4}) returns 10
// the function loops through a given slice and returns the sum of the slice elemts
func sumOfSlice(slice []int) int {
	sum := 0

	for _, calorie := range slice {
		sum += calorie
	}
	return sum
}

// TopThreeElves()
// the function calculates the total number of calories by the top three elves
func (q *Question) TopThreeElves() {
	totalCalories := []int{}

	elfFood := strings.Split(q.File, "\n\n")
	for _, food := range elfFood {
		currentFood := strings.Split(food, "\n")

		currentLargest := 0
		for _, calorie := range currentFood {
			if calorie == "" {
				continue
			}

			currentCalorie, err := strconv.Atoi(calorie)
			if err != nil {
				panic(err)
			}

			currentLargest += currentCalorie
		}

		totalCalories = append(totalCalories, currentLargest)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCalories)))
	q.PartTwo = sumOfSlice(totalCalories[0:3])
}
