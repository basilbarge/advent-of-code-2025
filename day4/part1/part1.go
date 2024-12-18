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
	//backwardDiagonal := make([]string, numDiagonals)


	regexp.MustCompile("XMAS|SAMX")

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

	for i := 0; i < numDiagonals; i++ {
		count := min(i, len(verticalLines) - 1)
		forwardDiagonal[i] = string(horizontalLines[min(i, len(horizontalLines) - 1)][max(0, i - len(verticalLines))])

		fmt.Printf("Iteration: %d\n", i)
		for count != 0 {
			y := min(count, len(horizontalLines) - 1) - 1
			x := min(max(0, count - len(verticalLines)) + 1 + i, len(verticalLines) - 1)
			fmt.Printf("Forward diag y: %d, Forward diag x: %d\n", y, x)
			forwardDiagonal[i] += string(horizontalLines[x][y])
			count--
		}
	}

	fmt.Println("Forward Diagonal")
	for _, forDiag := range forwardDiagonal {
		fmt.Println(forDiag)
	}
}
