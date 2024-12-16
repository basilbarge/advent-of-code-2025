package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/aoc2024/utils"
)

func main() {
	lines := utils.ReadLines("../input.txt")

	code := strings.Join(lines, "")

	validOperation := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	operations := validOperation.FindAll([]byte(code), -1)

	sum := 0
	for _, byte := range operations {
		operands := GetOperands(string(byte))

		sum += (operands[0] * operands[1])
	}

	fmt.Println(sum)
}

func GetOperands(operation string) [2]int {

	functionOperands := regexp.MustCompile(`(\d{1,3}),(\d{1,3})`)

	result := functionOperands.FindStringSubmatch(operation)

	operand1, err := strconv.Atoi(result[1])
	if err != nil {
		panic("Error parsing first operand")
	}

	operand2, err := strconv.Atoi(result[2])

	if err != nil {
		panic("Error parsing second operand")
	}

	return [2]int{operand1, operand2}
}
