package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

func main() {

	rows := utils.ReadLines("../input.txt")

	var safeCount int

	for _, row := range rows {
		levels := strings.Split(row, " ")

		var isSafe bool = true

		isIncreasingOrDecreasing := IsOnlyIncreasingOrDecreasing(levels)

		for i := 0; i + 1 < len(levels); i++ {
			next, err := strconv.Atoi(levels[i + 1])

			if err != nil {
				panic(fmt.Sprintf("Could not convert %s to int", levels[i + 1]))
			}

			current, err := strconv.Atoi(levels[i])

			if err != nil {
				panic(fmt.Sprintf("Could not convert %s to int", levels[i]))
			}


			isIncreasingInRange := current + 1 <= (next) && current + 3 >= (next)
			isDecreasingInRange := current - 1 >= (next) && current - 3 <= (next)

			// The next number differs by too much
			if (!(isIncreasingInRange || isDecreasingInRange)) {
				isSafe = false
				break
			}

		}

		if isSafe && isIncreasingOrDecreasing {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func IsOnlyIncreasingOrDecreasing(nums []string) bool {
	var isIncreasing bool
	var isDecreasing bool

	for idx := 0; idx + 1 < len(nums); idx++{
		next, err := strconv.Atoi(nums[idx + 1])

		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int", nums[idx + 1]))
		}

		current, err := strconv.Atoi(nums[idx])

		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int", nums[idx]))
		}

		if (next > current) {
			if (isDecreasing) { return false }

			isIncreasing = true
		}

		if (current > next) {
			if (isIncreasing) { return false }

			isDecreasing = true
		}
	}

	return true
}
