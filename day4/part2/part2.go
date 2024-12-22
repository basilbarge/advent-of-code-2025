package main

import (
	"fmt"

	util "github.com/aoc2024/utils"
)

func main() {
	lines := util.ReadLines("../input.txt")

	sum := 0

	for linesIdx, line := range lines {
		for lineIdx, char := range line {
			if char == 'A' {
				sum += checkForwardDiagonal(lines, linesIdx, lineIdx)
			}
		}
	}

	fmt.Println(sum)

}

func checkForwardDiagonal(lines []string, row int, column int) int {
	if OutOfBounds(row, column, len(lines) - 1, len(lines[row]) - 1) {
		return 0
	}

	word := "A"
	wordForwardDiag := string(lines[row-1][column-1]) + word + string(lines[row+1][column+1])
	wordBackwardsDiag := string(lines[row+1][column-1]) + word + string(lines[row-1][column+1])

	if (wordForwardDiag == "MAS" || wordForwardDiag == "SAM") && (wordBackwardsDiag == "MAS" || wordBackwardsDiag == "SAM") {
		return 1
	} else {
		return 0
	}
}

func OutOfBounds(row int, column int, maxRow int, maxColumn int) bool {
	return row < 1 || row > maxRow - 1 || column < 1 || column > maxColumn - 1
}
