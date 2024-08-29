package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
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
	fmt.Println(area(line))
}
// first step of the questions section
func area(line []string) int {
	var (
		surfacearea        int
		length             string
		width              string
		height             string
		splits             []string
		smallestarea       int
		totalwrappingpaper int
	)
	for _, ch := range line {
		splits = strings.Split(ch, "x")
		if len(splits) != 3 { 
			fmt.Println("Invalid line format:", ch)
			continue
		}
		length = splits[0]
		width = splits[1]
		height = splits[2]
		surfacearea = 2*atoi(length)*atoi(width) + 2*atoi(height)*atoi(width) + 2*atoi(length)*atoi(height)
		smallestarea = atoi(length) * atoi(width)
		area2 := atoi(length) * atoi(height)
		area3 := atoi(height) * atoi(width)
		if area2 < smallestarea {
			smallestarea = area2
		}
		if area3 < smallestarea {
			smallestarea = area3
		}
		totalwrappingpaper += (surfacearea + smallestarea)

	}
	return totalwrappingpaper

}
func atoi(a string) int {
	sign := 1
	if a[0] == '-' {
		sign = -1
		a = a[1:]
	}
	var num int
	for _, ch := range a {
		if ch < '0' || ch > '9' {
			return 0
		}
		num = num*10 + int(ch-'0')
	}
	if sign < 1 {
		num = num * -1
	}
	return num
}

// second step of the question
