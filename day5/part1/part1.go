package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

func main() {
	input := utils.ReadLines("../input.txt")

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

	sum := 0
	for _, pageOrder := range pages {
		pagesInOrder := strings.Split(pageOrder, ",")

		orderValid := PagesAreInOrder(pagesInOrder, rules)

		if orderValid {
			middlePage, err := strconv.Atoi(pagesInOrder[len(pagesInOrder)/2])
			if err != nil {
				panic("Could not convert middle page to integer")
			}

			sum += middlePage
		}
	}

	fmt.Println(sum)
}

func PagesAreInOrder(pagesInOrder []string, rules map[int][]int) bool {
	for currentPageIdx, page := range pagesInOrder {
		pageAsInt, err := strconv.Atoi(page)
		if err != nil {
			panic("Could not convert page to integer")
		}

		pagesAfter, exists := rules[pageAsInt]

		if !exists {
			continue
		}

		for i := currentPageIdx - 1; i >= 0; i-- {
			currentPageAsInt, err := strconv.Atoi(pagesInOrder[i])
			if err != nil {
				panic("Could not convert current page to integer")
			}

			if slices.Contains(pagesAfter, currentPageAsInt) {
				return false
			}
		}
	}

	return true
}
