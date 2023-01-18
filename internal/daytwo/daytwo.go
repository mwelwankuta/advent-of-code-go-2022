package daytwo

import (
	"strings"

	"github.com/mwelwankuta/advent-of-code-go-2022/internal/question"
)

type Question struct {
	question.Question
	match []string
}

const (
	rockBalance    = 1 // Balance for choosing rock
	paperBalance   = 2 // Balance for choosing paper
	scissorBalance = 3 // Balance for choosing scissors
)

const (
	playerRock     = "X" // Rock sign for player
	playerPaper    = "Y" // Paper sign for player
	playerScissors = "Z" // Scissors sign for player
)

const (
	oponentRock     = "A" // Rock sign for oponent
	oponentPaper    = "B" // Paper sign for oponent
	oponentScissors = "C" // Scissors sign for oponent
)

var signBalanceMap = map[string]int{
	// Oponent balance maps
	"A": rockBalance, // Rock sign map to value
	"X": rockBalance,
	"B": paperBalance, // Paper sign map to value

	// Player balance maps
	"Y": paperBalance,
	"C": scissorBalance, // Scissors sign map to value
	"Z": scissorBalance,
}

// getStandardizedSign(player string) int
// returns balance of passed sign
// the functions takes in a sign from either the player or oponent and returns
// the balance the player or oponent has for choosing that sign
func getStandardizedSign(player string) int {
	return signBalanceMap[player]
}

func (q *Question) loadMatches(match []string) {
	q.match = match
}

// CalculateRPSTotal()
// calculates the total the player would have after winning the rock, paper, scissors
// matches
func (q *Question) CalculateRPSTotal() {
	match := strings.Split(q.File, "\n")
	q.loadMatches(match)

	winBalance := 0 // Balance accumulated from getting points for match rounds

	for _, currentGame := range q.match {
		round := strings.Split(currentGame, " ")

		if len(round) < 2 {
			continue
		}

		player := round[1]  // Sign the player chooses
		oponent := round[0] // Sign the oponent chooses

		switch {
		case getStandardizedSign(player) == getStandardizedSign(oponent):
			winBalance += 3 // For drawing with oponent

			if player == playerScissors {
				winBalance += scissorBalance // For choosing scissors
			} else if player == playerRock {
				winBalance += rockBalance // For choosing rock
			} else if player == playerPaper {
				winBalance += paperBalance // For choosing paper
			}

		case player == playerRock:
			if oponent == oponentScissors {
				winBalance += (rockBalance + 6) // For choosing rock and winning
			} else {
				winBalance += rockBalance // For choosing rock
			}

		case player == playerPaper:
			if oponent == oponentRock {
				winBalance += (paperBalance + 6) // For choosing paper and winning
			} else {
				winBalance += paperBalance // For choosing paper
			}

		case player == playerScissors:
			if oponent == oponentPaper {
				winBalance += (scissorBalance + 6) // For choosing scissors and winning
			} else {
				winBalance += scissorBalance // For choosing scissors
			}
		}

	}

	// Set part two answer
	q.PartOne = winBalance
}

const (
	playerLose = "X" // Value of round outcome when player loses
	playerDraw = "Y" // Value of round outcome when player draws
	playerWin  = "Z" // Value of round outcome when player wins
)

// CalculateRoundOutcomes()
// calculates when to use which sign
func (q *Question) CalculateRoundOutcomes() {
	winBalance := 0 // Balance accumulated from getting points for match rounds

	for _, currentGame := range q.match {
		round := strings.Split(currentGame, " ")

		if len(round) < 2 {
			continue
		}

		oponent := round[0]      // Sign the component chooses
		howShouldEnd := round[1] // How round needs to end

		switch {
		case oponent == oponentRock: // when the oponent chooses rock
			if howShouldEnd == playerDraw {
				winBalance += (rockBalance + 3) // For choosing rock and drawing
			} else if howShouldEnd == playerLose {
				winBalance += (scissorBalance) // For choosing scissors
			} else if howShouldEnd == playerWin {
				winBalance += (paperBalance + 6) // For choosing paper and winning
			}

		case oponent == oponentPaper: // when the oponent chooses paper
			if howShouldEnd == playerDraw {
				winBalance += (paperBalance + 3) // For choosing paper and drawing
			} else if howShouldEnd == playerLose {
				winBalance += (rockBalance) // For choosing rock
			} else if howShouldEnd == playerWin {
				winBalance += (scissorBalance + 6) // For choosing scissors and winning
			}

		case oponent == oponentScissors: // when the oponent chooses scissors
			if howShouldEnd == playerDraw {
				winBalance += (scissorBalance + 3) // For choosing scissor and drawing
			} else if howShouldEnd == playerLose {
				winBalance += (paperBalance) // For choosing paper
			} else if howShouldEnd == playerWin {
				winBalance += (rockBalance + 6) // For choosing rock and winning
			}

		}

	}

	// Set part two answer
	q.PartTwo = winBalance
}
