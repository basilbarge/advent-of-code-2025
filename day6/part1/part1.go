package main

import (
	"errors"
	"fmt"

	util "github.com/aoc2024/utils"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
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
		guard.Location.y--
	} else if guard.Dir == Down {
		guard.Location.y++
	} else if guard.Dir == Right {
		guard.Location.x++
	} else { //Guard direction must be left
		guard.Location.x--
	}
}

func (guard *Guard) Turn() {
	guard.Dir = (guard.Dir + 1) % 4
}

func (guard *Guard) BackUp() {
	if guard.Dir == Up {
		guard.Location.y++
	} else if guard.Dir == Down {
		guard.Location.y--
	} else if guard.Dir == Right {
		guard.Location.x--
	} else { //Guard direction must be left
		guard.Location.x++
	}
}

func main() {
	input := util.ReadLines("../input.txt")

	var guard Guard

	guard.FindGuard(input)
	guard.DetermineDirection(input)

	//Mark where the guard is now
	input[guard.Location.y] = input[guard.Location.y][:guard.Location.x] + "X" + input[guard.Location.y][guard.Location.x+1:]

	moveGuard := true

	for moveGuard {
		guard.Move()
		err := UpdateRoom(input, &guard)

		// The guard has left the room
		if err != nil {
			fmt.Println(err)
			moveGuard = false
		}
	}

	visitedSpaces := 0
	for _, line := range input {
		for _, char := range line {
			if char == 'X' {
				visitedSpaces++
			}
		}
	}

	fmt.Printf("The guard visited %d spaces\n", visitedSpaces)
}

func UpdateRoom(room []string, guard *Guard) error {
	// Guard has left the room
	if guard.Location.y < 0 || guard.Location.y > len(room)-1 || guard.Location.x < 0 || guard.Location.x > len(room[0])-1 {
		return errors.New("The guard is now outside the room")
	}

	// Guard has run into an obstacle
	if room[guard.Location.y][guard.Location.x] == '#' {
		guard.BackUp()
		guard.Turn()
		guard.Move()
	}

	room[guard.Location.y] = room[guard.Location.y][:guard.Location.x] + "X" + room[guard.Location.y][guard.Location.x+1:]

	return nil
}
