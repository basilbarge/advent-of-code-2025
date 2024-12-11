package main

import (
	"fmt"
	"math"
	"path/filepath"
	"sort"
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

	sort.Ints(left)
	sort.Ints(right)

	sums := make([]float64, len(left))

	for idx, num := range left {
		sums[idx] = math.Abs(float64(num - right[idx]))
	}

	var sum float64 = 0

	for _, values := range sums {
		sum += values
	}

	fmt.Printf("%f\n", sum)
}


