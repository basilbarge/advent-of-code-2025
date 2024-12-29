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
	CurrentLocation Point
	StartLocation   Point
	Dir             Direction
	Stuck           bool
	VisitedSpaces   map[Point]Direction
}

func (guard *Guard) DetermineDirection(grid []string) {
	guardSymbol := rune(grid[guard.CurrentLocation.y][guard.CurrentLocation.x])

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
				guard.CurrentLocation.x = colIdx
				guard.CurrentLocation.y = rowIdx
				guard.StartLocation.x = colIdx
				guard.StartLocation.y = rowIdx
				return
			}
		}
	}

	panic("There was no guard found in the room")
}

func (guard *Guard) Move() {
	if guard.Dir == Up {
		guard.CurrentLocation.y--
	} else if guard.Dir == Down {
		guard.CurrentLocation.y++
	} else if guard.Dir == Right {
		guard.CurrentLocation.x++
	} else { //Guard direction must be left
		guard.CurrentLocation.x--
	}

}

func (guard *Guard) Turn() {
	guard.Dir = (guard.Dir + 1) % 4
}

func (guard *Guard) BackUp() {
	if guard.Dir == Up {
		guard.CurrentLocation.y++
	} else if guard.Dir == Down {
		guard.CurrentLocation.y--
	} else if guard.Dir == Right {
		guard.CurrentLocation.x--
	} else { //Guard direction must be left
		guard.CurrentLocation.x++
	}
}

func (guard *Guard) IsStuck() bool {
	currentDir, pres := guard.VisitedSpaces[guard.CurrentLocation]

	// Guard has been to this spot before and is going the
	// the same direction as they were the last time they were
	// in this spot --> guard is stuck in loop
	if pres && currentDir == guard.Dir {
		return true
	}

	return false
}

func main() {
	input := util.ReadLines("../input.txt")
	original_input := make([]string, len(input))

	copy(original_input, input)
	//util.PrintInput(original_input)

	guard := Guard{
		VisitedSpaces: make(map[Point]Direction),
	}

	guard.FindGuard(input)
	guard.DetermineDirection(input)

	//Mark where the guard is now
	input[guard.CurrentLocation.y] = input[guard.CurrentLocation.y][:guard.CurrentLocation.x] + "X" + input[guard.CurrentLocation.y][guard.CurrentLocation.x+1:]

	moveGuard := true

	for moveGuard {
		guard.Move()
		_, err := UpdateRoom(input, &guard)

		// The guard has left the room
		if err != nil {
			fmt.Println(err)
			moveGuard = false
		}
	}

	potentialObstacles := make(map[Point]Direction)

	for k, v := range guard.VisitedSpaces {
		potentialObstacles[k] = v
	}

	stuckCount := 0
	loop := 1
	//Mark all potential spaces for a new obstacle for debugging
	for location := range potentialObstacles {
		//if loop != 0 {
		//	fmt.Println("Skipping")
		//	continue
		//}

		//loop++

		//input[location.y] = input[location.y][:location.x] + "O" + input[location.y][location.x+1:]
		// Make a new guard to run test scenarios
		testGuard := Guard{
			VisitedSpaces: make(map[Point]Direction),
		}

		// Reset the room to the original way it was
		copy(input, original_input)

		testGuard.FindGuard(input)
		testGuard.DetermineDirection(input)

		fmt.Printf("Testing location %v. This is %d/%d\n", location, loop, len(potentialObstacles))
		// Add new obstacle at visited location
		input[location.y] = input[location.y][:location.x] + "#" + input[location.y][location.x+1:]
		//fmt.Println("Room with current obstacle looks like")
		//util.PrintInput(input)

		moveGuard := true

		for moveGuard {
			testGuard.Move()
			stuck, err := UpdateRoom(input, &testGuard)
			//fmt.Println("After move")
			//fmt.Printf("Current location is %d, %d\n", testGuard.CurrentLocation.x, testGuard.CurrentLocation.y)
			//util.PrintInput(input)

			if stuck {
				//fmt.Println("The guard is stuck")
				stuckCount++
				moveGuard = false
				loop++
				//util.PrintInput(input)
			}

			// The guard has left the room
			if err != nil {
				//fmt.Println(err)
				moveGuard = false
				loop++
			}
		}

		fmt.Printf("So far, the guard has gotten stuck in %d scenarios\n", stuckCount)
	}

	//util.PrintInput(input)

	fmt.Printf("The guard got stuck in %d scenarios\n", stuckCount)
}

func UpdateRoom(room []string, guard *Guard) (bool, error) {
	// Guard has left the room
	if guard.CurrentLocation.y < 0 || guard.CurrentLocation.y > len(room)-1 || guard.CurrentLocation.x < 0 || guard.CurrentLocation.x > len(room[0])-1 {
		return false, errors.New("The guard is now outside the room")
	}

	if room[guard.CurrentLocation.y][guard.CurrentLocation.x] == '#' {
		guard.BackUp()
		guard.Turn()
		guard.Move()
		if room[guard.CurrentLocation.y][guard.CurrentLocation.x] == '#' {
			guard.BackUp()
			guard.Turn()
			guard.Move()
		}
	}

	room[guard.CurrentLocation.y] = room[guard.CurrentLocation.y][:guard.CurrentLocation.x] + "*" + room[guard.CurrentLocation.y][guard.CurrentLocation.x+1:]

	if guard.IsStuck() {
		//fmt.Println("Guard is stuck")
		return true, nil
	}

	currentPoint := Point{guard.CurrentLocation.x, guard.CurrentLocation.y}
	_, pointVisited := guard.VisitedSpaces[currentPoint]

	if currentPoint != guard.StartLocation && !pointVisited {
		guard.VisitedSpaces[currentPoint] = guard.Dir
	}
	//room[guard.CurrentLocation.y] = room[guard.CurrentLocation.y][:guard.CurrentLocation.x] + "X" + room[guard.CurrentLocation.y][guard.CurrentLocation.x+1:]

	return false, nil
}
