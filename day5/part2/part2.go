package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc2024/utils"
)

type PageOrderError struct {
	msg             string
	pageOutOfOrder  int
	pageToComeAfter int
}

func (e *PageOrderError) Error() string {
	return e.msg
}

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

		_, err := PagesAreInOrder(pagesInOrder, rules)

		if err != nil {
			SwapPages(pagesInOrder, err.pageOutOfOrder, err.pageToComeAfter)

			for i := 0; i < len(pagesInOrder)*50; i++ {
				_, orderErr := PagesAreInOrder(pagesInOrder, rules)

				if orderErr != nil {
					SwapPages(pagesInOrder, orderErr.pageOutOfOrder, orderErr.pageToComeAfter)
					continue
				}

				break
			}

			middlePage, err := strconv.Atoi(pagesInOrder[len(pagesInOrder)/2])
			if err != nil {
				panic("Could not convert middle page to integer")
			}
			sum += middlePage
		}
	}

	fmt.Println(sum)
}

func PagesAreInOrder(pagesInOrder []string, rules map[int][]int) (bool, *PageOrderError) {
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
				err := &PageOrderError{
					fmt.Sprintf("Pages are out of order. %d should come before %d", pageAsInt, currentPageAsInt),
					currentPageIdx,
					i,
				}

				return false, err
			}
		}
	}

	return true, nil
}

func SwapPages(pages []string, pageToSwap1, pageToSwap2 int) {
	pages[pageToSwap1], pages[pageToSwap2] = pages[pageToSwap2], pages[pageToSwap1]
}

func Factorial(input int) int {
	if input == 0 {
		return 1
	}

	return input * Factorial(input - 1)
}
