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

func Reverse(str string) string {
	var rts string
	for i := len(str) - 1; i >= 0; i-- {
		rts += string(str[i])
	}
	return rts
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
