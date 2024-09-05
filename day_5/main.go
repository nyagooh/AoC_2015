package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Println(IsNice(line, NiceString))
	fmt.Println(IsNice(line, NiceStr))
}

func NiceString(s string) bool {
	var (
		count        int
		result       bool
		founddoubles bool
	)
	vowels := "aeiou"
	for _, char := range s {
		for _, val := range vowels {
			if string(char) == string(val) {
				count++
			}
		}
	}

	if count < 3 {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			founddoubles = true
			result = true
			break
		}
	}
	if !founddoubles {
		result = false
		return false
	}
	//checks for invalid string
	//substring algorithim
	invalidStrings := []string{"ab", "cd", "pq", "xy"}
	for _, invalid := range invalidStrings {
		lenInvalid := len(invalid)
		for i := 0; i <= len(s)-lenInvalid; i++ {
			if s[i:i+lenInvalid] == invalid {
				result = false
				break
			}
		}
	}
	return result
}

type Nicechecker func(string) bool

func IsNice(s []string, checker Nicechecker) int {
	count := 0
	for _, str := range s {
		if checker(str) {
			count++
		}
	}
	return count
}

// step 2 question
func NiceStr(s string) bool {
	var result bool
	var ok bool
	for i := 0; i < len(s)-1; i++ {
		pairs := s[i : i+2]
		if strings.Contains(s[i+2:], pairs) {
			result = true
		}
	}
	if result {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			ok = true
		}
	}
}
return ok
}
