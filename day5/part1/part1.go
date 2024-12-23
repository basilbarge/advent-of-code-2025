package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

func main() {
	input := utils.ReadLines("../sample.txt")

	var ruleset []string
	var pages []string

	markPages := false

	for _, line := range input {
		if line == "" {
			markPages = true
			continue
		}

		if markPages {
			pages = append(pages, line)
		} else {
			ruleset = append(ruleset, line)
		}
	}

	rules := make(map[int][]int)

	for _, rule := range ruleset {
		assignment := strings.Split(rule, "|")

		key, err := strconv.Atoi(assignment[0])
		if err != nil {
			panic("Could not convert key to integer")
		}

		value, err := strconv.Atoi(assignment[1])
		if err != nil {
			panic("Could not convert value to integer")
		}

		rules[key] = append(rules[key], value)
	}


	fmt.Printf("Rules\n%d\n", rules)
	fmt.Printf("Pages\n%s\n", pages)
}
