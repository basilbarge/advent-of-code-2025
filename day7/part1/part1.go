package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

type Operator bool

const (
	Addition Operator = false
	Multiply Operator = true
)

func main() {

	input := utils.ReadLines("../sample.txt")

	problems := make(map[int64][]int64)

	for _, line := range input {
		splitProblem := strings.Split(line, ":")
		target, err := strconv.ParseInt(splitProblem[0], 10, 32)

		if err != nil {
			fmt.Printf("There was a problem parsing the target as an int. %s\n", err)
		}

		operands := strings.Split(strings.Trim(splitProblem[1], " "), " ")

		for _, operand := range operands {
			opAsInt, err := strconv.ParseInt(operand, 10, 8)

			if err != nil {
				fmt.Printf("There was a problem parsing %s as an int. %s\n", operand, err)
			}

			problems[target] = append(problems[target], opAsInt)
		}
	}

	fmt.Println(problems)
}
