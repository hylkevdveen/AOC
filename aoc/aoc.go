package aoc

import (
	"bufio"
	"log"
	"os"
)

func InputLines() []string {
	ret := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("read input.txt: %w", err)
	}
	return ret
}

// Reverse returns the given string in reverse
func Reverse(str string) string {
	var rts string
	for i := len(str) - 1; i >= 0; i-- {
		rts += string(str[i])
	}
	return rts
}

// AbsDiffInt returns the absolute difference between two given integers
func AbsDiffInt(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// GCD finds the greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

var DigitStrings = []string{
	"zero",
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
