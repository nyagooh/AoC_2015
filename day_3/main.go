package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 || len(args) < 1 {
		return
	}
	filename, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	fmt.Println(countHouses(string(filename)))
	fmt.Println(TurnHouses(string(filename)))

}

func countHouses(directions string) int {
	x, y := 0, 0
	houses := make(map[string]bool)
	houses[fmt.Sprintf("%d,%d", x, y)] = true

	for _, dir := range directions {
		switch dir {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		houses[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return len(houses)
}

// second question
func TurnHouses(directions string) int {
	santax, santay := 0, 0
	robotx, roboty := 0, 0
	houses := make(map[string]bool)
	//starting point
	houses["0,0"] = true
	// to know which turn to take
	for i, dir := range directions {
		var x, y *int
		if i%2 == 0 {
			x, y = &santax, &santay
		} else {
			x, y = &robotx, &roboty
		}
		switch dir {
		case '^':
			*y++
		case 'v':
			*y--
		case '>':
			*x++
		case '<':
			*x--
		}
		houses[fmt.Sprintf("%d,%d", *x, *y)] = true
	}

	return len(houses)
}
