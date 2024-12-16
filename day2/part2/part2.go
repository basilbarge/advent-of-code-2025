package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

func main() {

	rows := utils.ReadLines("../input.txt")

	safeCount := 0
	for _, row := range rows {
		level := strings.Split(row, " ")

		options := make([][]string, len(level) + 1)

		options[0] = level
		for idx := range level {
			option := make([]string, idx, len(level))
			copy(option, level[:idx])
			options[idx + 1] = append(option, level[idx + 1:]...)
		}

		hasSafeOption := false
		for _, option := range options {
			hasSafeOption = IsSafe(option)

			if hasSafeOption { 
				safeCount++
				break
			}
		}
	}

	fmt.Println(safeCount)
}

func IsSafe(option []string) bool {
	var isSafe bool = true

	isIncreasingOrDecreasing := IsOnlyIncreasingOrDecreasing(option)

	for i := 0; i + 1 < len(option); i++ {
		next, err := strconv.Atoi(option[i + 1])

		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int", option[i + 1]))
		}

		current, err := strconv.Atoi(option[i])

		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int", option[i]))
		}


		isIncreasingInRange := current + 1 <= (next) && current + 3 >= (next)
		isDecreasingInRange := current - 1 >= (next) && current - 3 <= (next)

		// The next number differs by too much
		if (!(isIncreasingInRange || isDecreasingInRange)) {
			isSafe = false
			break
		}

	}

	return isSafe && isIncreasingOrDecreasing
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
