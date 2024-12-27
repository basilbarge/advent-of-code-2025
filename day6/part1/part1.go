package main

import (
	"fmt"

	util "github.com/aoc2024/utils"
)

type Direction int

const (
	Up Direction = iota
	Right
	Left
	Down
)

type test struct {
	x int
	y int
}

type Point struct {
	x int
	y int
}

type Guard struct {
	Location Point
	Dir      Direction
}

func (guard *Guard) DetermineDirection(grid []string) {
	guardSymbol := rune(grid[guard.Location.y][guard.Location.x])

	if guardSymbol == '<' {
		guard.Dir = Left
	} else if guardSymbol == '>' {
		guard.Dir = Right
	} else if guardSymbol == '^' {
		guard.Dir = Up
	} else { //symbol must be v in this case
		guard.Dir = Down
	}
}

func (guard *Guard) FindGuard(room []string) {
	for rowIdx, line := range room {
		for colIdx, char := range line {
			if char == '<' || char == 'v' || char == '>' || char == '^' {
				guard.Location.x = colIdx
				guard.Location.y = rowIdx
				return
			}
		}
	}

	panic("There was no guard found in the room")
}

func (guard *Guard) Move() {
	if guard.Dir == Up {
		guard.Location.y++
	} else if guard.Dir == Down {
		guard.Location.y--
	} else if guard.Dir == Right {
		guard.Location.x++
	} else { //Guard direction must be left
		guard.Location.x--
	}
}

func main() {
	input := util.ReadLines("../sample.txt")

	var guard Guard

	guard.FindGuard(input)
	guard.DetermineDirection(input)

	//Mark where the guard is now
	input[guard.Location.y] = input[guard.Location.y][:guard.Location.x] + "X" + input[guard.Location.y][guard.Location.x+1:]


	guard.Move()
	UpdateRoom(input, guard)

	for _, line := range input {
		fmt.Println(line)
	}
}

func UpdateRoom (room []string, guard Guard) {
	room[guard.Location.y] = room[guard.Location.y][:guard.Location.x] + "X" + room[guard.Location.y][guard.Location.x+1:]
}
