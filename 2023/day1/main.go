package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var nums = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	sum := 0
	for _, line := range aoc.InputLines() {
		sum += Part2(line)
	}
	fmt.Println(sum)
}

func Part1(line string) int {
	number := 0
	str := []rune(line)
	fmt.Printf("%s\n", line)
	// First digit
	for i := 0; i < len(str); i++ {
		if unicode.IsDigit(str[i]) {
			digit, _ := strconv.Atoi(string(str[i]))
			fmt.Printf("First digit: %d\n", digit)
			number += 10 * digit
			break
		}
	}
	// Last digit
	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsDigit(str[i]) {
			digit, _ := strconv.Atoi(string(str[i]))
			fmt.Printf("Last digit: %d\n", digit)
			number += digit
			break
		}
	}
	return number
}

func Part2(line string) int {
	number := 0
	ucLine := []rune(line)
	fmt.Printf("%s\n", line)
	// First digit
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(ucLine[i]) {
			digit, _ := strconv.Atoi(string(ucLine[i]))
			fmt.Printf("First digit: %d\n", digit)
			number += 10 * digit
			break
		}
		found := false
		for val, digit := range nums {
			if strings.HasPrefix(line[i:], digit) {
				fmt.Printf("First digit: %d\n", val+1)
				number += 10 * (val + 1)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	// Last digit
	for i := len(line); i > 0; i-- {
		if unicode.IsDigit(ucLine[i-1]) {
			digit, _ := strconv.Atoi(string(ucLine[i-1]))
			fmt.Printf("Last digit: %d\n", digit)
			number += digit
			break
		}
		found := false
		for val, digit := range nums {
			if strings.HasSuffix(line[:i], digit) {
				fmt.Printf("Last digit: %d\n", val+1)
				number += val + 1
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	return number
}
