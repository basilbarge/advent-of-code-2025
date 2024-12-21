package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/aoc2024/utils"
)


func main() {

	input := utils.ReadLines("../input.txt")

	code := ""
	for _, line := range input {
		code += line
	}


	functions := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	commands := functions.FindAll([]byte(code), -1)
	sum := 0

	execute := true
	for _, command := range commands {

		if string(command) == "do()" {
			execute = true
		} else if string(command) == "don't()" {
			execute = false
		} else if execute {
			sum += MultArgs(string(command))
		}
	}

	fmt.Println(sum)

}

func MultArgs(multCommand string) int {
	operandMatch := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	operands := operandMatch.FindAllSubmatch([]byte(multCommand), -1)[0]

	operand1, err := strconv.Atoi(string(operands[1]))
	if err != nil {
		fmt.Printf("There was an error parsing the first argument in: %s", multCommand)
	}

	operand2, err := strconv.Atoi(string(operands[2]))
	if err != nil {
		fmt.Printf("There was an error parsing the second argument in: %s", multCommand)
	}

	return operand1 * operand2
}
