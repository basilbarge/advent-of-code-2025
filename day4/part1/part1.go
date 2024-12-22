// Answer should be 2591
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
			if char == 'X' {
				sum += checkHorizontal(lines, linesIdx, lineIdx)
				sum += checkHorizontalBackwards(lines, linesIdx, lineIdx)
				sum += checkVertical(lines, linesIdx, lineIdx)
				sum += checkVerticalBackwards(lines, linesIdx, lineIdx)
				sum += checkForwardDiagonal(lines, linesIdx, lineIdx)
				sum += checkForwardDiagonalBackwards(lines, linesIdx, lineIdx)
				sum += checkBackDiagonal(lines, linesIdx, lineIdx)
				sum += checkBackDiagonalBackwards(lines, linesIdx, lineIdx)
			}
		}
	}

	fmt.Println(sum)

}

func checkHorizontal(lines []string, row int, column int) int {
	if column > len(lines[row])-4 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row][column+1]) + string(lines[row][column+2]) + string(lines[row][column+3])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkHorizontalBackwards(lines []string, row int, column int) int {
	if column < 3 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row][column-1]) + string(lines[row][column-2]) + string(lines[row][column-3])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkVertical(lines []string, row int, column int) int {
	if row > len(lines)-4 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row+1][column]) + string(lines[row+2][column]) + string(lines[row+3][column])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkVerticalBackwards(lines []string, row int, column int) int {
	if row < 3 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row-1][column]) + string(lines[row-2][column]) + string(lines[row-3][column])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkForwardDiagonal(lines []string, row int, column int) int {
	if row < 3 || column > len(lines[row])-4 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row-1][column+1]) + string(lines[row-2][column+2]) + string(lines[row-3][column+3])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkForwardDiagonalBackwards(lines []string, row int, column int) int {
	if row > len(lines)-4 || column < 3 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row+1][column-1]) + string(lines[row+2][column-2]) + string(lines[row+3][column-3])

	madeWord := word + string(nextLetters)


	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkBackDiagonal(lines []string, row int, column int) int {
	if row > len(lines)-4 || column > len(lines[row])-4 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row+1][column+1]) + string(lines[row+2][column+2]) + string(lines[row+3][column+3])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}

func checkBackDiagonalBackwards(lines []string, row int, column int) int {
	if row < 3 || column < 3 {
		return 0
	}

	word := "X"
	nextLetters := string(lines[row-1][column-1]) + string(lines[row-2][column-2]) + string(lines[row-3][column-3])

	madeWord := word + string(nextLetters)

	if madeWord == "XMAS" {
		//fmt.Printf("row %d, column %d: %s\n", row, column, madeWord)
		return 1
	} else {
		return 0
	}
}
