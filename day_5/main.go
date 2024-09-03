package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	fmt.Println(IsNice(line))
}

func NiceString(s string) bool {
	var (
		count        int
		founddoubles bool
	)
	vowels := "aeiou"
	//checks for vowels

	for _, char := range s {
		for _, val := range vowels {
			if string(char) == string(val) {
				count++
			}
		}
	}

	//checks for doubles
	if count < 3 {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			founddoubles = true
			break
		}
	}
	if !founddoubles {
		return false
	}

	//checks for invalid string
	invalidStrings := []string{"ab", "cd", "pq", "xy"}
	for _, invalid := range invalidStrings {
		lenInvalid := len(invalid)
		for i := 0; i <= len(s)-lenInvalid; i++ {
			match := false
			for j := 0; j < lenInvalid; j++ {
				if s[i+j] == invalid[j] {
					match = true
					break
				}
			}
			if match {
				return false
			}
		}

	}
	return true
}
func IsNice(s []string) int {
	count := 0
	for _,str :=range s{
		if NiceString(str){
			count++
		}
	}
	return count
}
