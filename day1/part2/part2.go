package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

func main() {

	data := utils.ReadLines(filepath.Join("..", "input.txt"))

	left := make([]int, len(data))
	right := make([]int, len(data))

	for idx, str := range data {
		rows := strings.Split(str, "   ")

		var errLeft error
		var errRight error
		left[idx], errLeft = strconv.Atoi(rows[0])
		right[idx], errRight = strconv.Atoi(rows[1])

		if (errLeft != nil) {
			panic(fmt.Sprintf("Error parsing left side integer: %s", errLeft))
		}
		if (errRight != nil) {
			panic(fmt.Sprintf("Error parsing right side integer: %s", errRight))
		}
	}

	count := make(map[float64]int)
	for _, left_num := range left {

		for _, right_num := range right {
			if (right_num == left_num) {
				count[float64(left_num)] += 1
			}
		}
	}

	var sum float64 = 0

	for key, value := range count {
		sum += key * float64(value)
	}

	fmt.Printf("%f\n", sum)
}
