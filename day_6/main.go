package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const size = 1000

func main() {
	args := os.Args[1:]
	if len(args) > 1 || len(args) < 1 {
		return
	}
	filename, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(filename)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	// Initialize a 1000x1000 grid of lights, all initially off (false)
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}

	// Process each command
	for _, command := range line{
		processCommand(grid, command)
	}

	// Count the number of lights that are on
	litCount := countLitLights(grid)
	fmt.Println("Number of lights lit:", litCount)
}

// Function to process commands
func processCommand(grid [][]bool, command string) {
	parts := strings.Fields(command)
	var action string
	var startX, startY, endX, endY int

	// Determine the action (turn on, turn off, toggle) and parse coordinates
	if parts[0] == "turn" {
		action = parts[1] // "on" or "off"
		startX, startY, endX, endY = parseCoordinates(parts[2], parts[4])
	} else if parts[0] == "toggle" {
		action = "toggle"
		startX, startY, endX, endY = parseCoordinates(parts[1], parts[3])
	}

	// Update the grid based on the action
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			if action == "on" {
				grid[i][j] = true // Turn on the light
			} else if action == "off" {
				grid[i][j] = false // Turn off the light
			} else if action == "toggle" {
				grid[i][j] = !grid[i][j] // Toggle the light
			}
		}
	}
}

// Function to parse coordinates from the command
func parseCoordinates(start, end string) (int, int, int, int) {
	var startX, startY, endX, endY int
	fmt.Sscanf(start, "%d,%d", &startX, &startY)
	fmt.Sscanf(end, "%d,%d", &endX, &endY)
	return startX, startY, endX, endY
}

// Function to count the number of lit lights
func countLitLights(grid [][]bool) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] { // If the light is on
				count++
			}
		}
	}
	return count
}
