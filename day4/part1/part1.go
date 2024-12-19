package main

import (
	"fmt"
	"regexp"

	util "github.com/aoc2024/utils"
)

func main() {
	horizontalLines := util.ReadLines("../input.txt")
	verticalLines := make([]string, len(horizontalLines[0]))

	numDiagonals := (len(horizontalLines) + len(horizontalLines[0]) - 1)

	forwardDiagonal := make([]string, numDiagonals)
	backwardDiagonal := make([]string, numDiagonals)


	for _, line := range horizontalLines {
		for idx, char := range line {
			verticalLines[idx] += string(char)
		}
	}

	for idx := 0; idx < len(horizontalLines)+len(verticalLines)-1; idx++ {
		xcount := 0
		if idx > len(verticalLines) {
			xcount = idx - len(verticalLines)
		}
		ycount := min(idx, len(horizontalLines)-1)

		forwardDiagonal[idx] = string(horizontalLines[ycount][xcount])

		xcount += 1
		ycount -= 1

		for (xcount < len(verticalLines)) && (ycount >= 0) {
			forwardDiagonal[idx] += string(horizontalLines[ycount][xcount])
			xcount++
			ycount--
		}
	}

	for idx := 0; idx < len(horizontalLines)+len(verticalLines)-1; idx++ {
		xcount := len(verticalLines) - 1
		if idx > len(horizontalLines)-1 {
			xcount = xcount - (idx - len(horizontalLines) + 1)
		}
		ycount := min(idx, len(horizontalLines)-1)

		backwardDiagonal[idx] = string(horizontalLines[ycount][xcount])

		xcount -= 1
		ycount -= 1

		for (ycount >= 0) && (xcount >= 0) {
			backwardDiagonal[idx] += string(horizontalLines[ycount][xcount])
			xcount--
			ycount--
		}
	}

	var matches []string

	matches = append(matches, FindMatches(horizontalLines)...)
	matches = append(matches, FindMatches(verticalLines)...)
	matches = append(matches, FindMatches(forwardDiagonal)...)
	matches = append(matches, FindMatches(backwardDiagonal)...)


	fmt.Println(len(matches))
}

func FindMatches(lines []string) []string {
	var matches []string 

	forKeyWord := regexp.MustCompile(`XMAS`)
	backKeyWord := regexp.MustCompile(`SAMX`)

	for _, line := range lines {
		forMatch := forKeyWord.FindAll([]byte(line), -1)
		backMatch := backKeyWord.FindAll([]byte(line), -1)

		for _, mat := range forMatch {
			matches = append(matches, string(mat))
		}

		for _, mat := range backMatch {
			matches = append(matches, string(mat))
		}
	}

	return matches
}
