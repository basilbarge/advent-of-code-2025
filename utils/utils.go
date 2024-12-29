package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) []string {

	data, err := os.ReadFile(path)

	if (err != nil) { panic(err) }

	data_str := string(data)

	data_str = strings.TrimSpace(data_str)

	return strings.Split(data_str, "\n")
}

func PrintInput(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}
