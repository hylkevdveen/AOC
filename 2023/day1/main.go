package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}

// Part1 Find the first and last digit of each string and sum the number you get from first+last
func Part1() int {
	sum := 0
	for _, line := range aoc.InputLines() {
		str := []rune(line)
		fmt.Printf("%s\n", line)
		// First digit
		for i := 0; i < len(str); i++ {
			if unicode.IsDigit(str[i]) {
				digit, _ := strconv.Atoi(string(str[i]))
				fmt.Printf("First digit: %d\n", digit)
				sum += 10 * digit
				break
			}
		}
		// Last digit
		for i := len(str) - 1; i >= 0; i-- {
			if unicode.IsDigit(str[i]) {
				digit, _ := strconv.Atoi(string(str[i]))
				fmt.Printf("Last digit: %d\n", digit)
				sum += digit
				break
			}
		}
	}
	return sum
}

// Part2 Do the same as Part1 except digits can be written out
func Part2() int {
	number := 0
	for _, line := range aoc.InputLines() {
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
			for val, digit := range aoc.DigitStrings {
				if strings.HasPrefix(line[i:], digit) {
					fmt.Printf("First digit: %d\n", val)
					number += 10 * (val)
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
			for val, digit := range aoc.DigitStrings {
				if strings.HasSuffix(line[:i], digit) {
					fmt.Printf("Last digit: %d\n", val)
					number += val
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return number
}
