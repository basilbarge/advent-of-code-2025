package main

import (
	"fmt"
	"regexp"

	util "github.com/aoc2024/utils"
)

func main() {
	horizontalLines := util.ReadLines("../sample.txt")
	verticalLines := make([]string, len(horizontalLines[0]))

	numDiagonals := (len(horizontalLines) + len(horizontalLines[0]) - 1)

	forwardDiagonal := make([]string, numDiagonals)
	backwardDiagonal := make([]string, numDiagonals)

	keyWord := regexp.MustCompile("XMAS|SAMX")

	fmt.Println("Horizontal")
	for _, line := range horizontalLines {

		for idx, char := range line {
			verticalLines[idx] += string(char)
		}
		fmt.Println(line)
	}

	fmt.Println("Vertical")
	for _, vertLine := range verticalLines {
		fmt.Println(vertLine)
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

	fmt.Println("Forward Diagonal")
	for _, forDiag := range forwardDiagonal {
		fmt.Println(forDiag)
	}

	for idx := 0; idx < len(horizontalLines)+len(verticalLines)-1; idx++ {
		xcount := len(verticalLines) - 1
		if idx > len(horizontalLines) - 1 {
			xcount = xcount - (idx - len(horizontalLines) + 1)
		}
		ycount := min(idx, len(horizontalLines) - 1)

		backwardDiagonal[idx] = string(horizontalLines[ycount][xcount])

		xcount -= 1
		ycount -= 1

		for (ycount >= 0) && (xcount >= 0) {
			backwardDiagonal[idx] += string(horizontalLines[ycount][xcount])
			xcount--
			ycount--
		}
	}

	fmt.Println("Backward Diagonal")
	for _, backDiag := range backwardDiagonal {
		fmt.Println(backDiag)
	}

	var matches []string

	for _, horizontal := range horizontalLines {
		match := keyWord.FindAll([]byte(horizontal), -1)	

		for _, mat := range match {
			matches = append(matches, string(mat))
		}
	}

	for _, vertical := range verticalLines {
		match := keyWord.FindAll([]byte(vertical), -1)	

		for _, mat := range match {
			matches = append(matches, string(mat))
		}
	}


	for _, forward := range forwardDiagonal {
		match := keyWord.FindAll([]byte(forward), -1)	

		for _, mat := range match {
			matches = append(matches, string(mat))
		}
	}

	for _, backward := range backwardDiagonal {
		match := keyWord.FindAll([]byte(backward), -1)	

		for _, mat := range match {
			matches = append(matches, string(mat))
		}
	}

	fmt.Println(len(matches))
}
