package manual

import (
	"../../../aoc"
	"strconv"
	"strings"
	"unicode"
)

type Manual struct {
	Seeds    []int
	Mappings [][]*Mapping
}

type Mapping struct {
	TargetStart int
	SourceStart int
	Range       int
}

func ReadManual() *Manual {
	lines := aoc.InputLines()
	seeds := ReadSeeds(lines[0])
	var mappings [][]*Mapping
	lines = lines[3:]
	var mapping []*Mapping
	for _, line := range lines {
		if strings.IndexFunc(line, func(r rune) bool {
			return unicode.IsDigit(r)
		}) == 0 {
			mapping = append(mapping, ReadMapping(line))
		} else {
			if mapping != nil {
				mappings = append(mappings, mapping)
			}
			mapping = nil
		}
	}
	mappings = append(mappings, mapping)
	return &Manual{
		Seeds:    seeds,
		Mappings: mappings,
	}
}

func ReadMapping(line string) *Mapping {
	values := strings.Split(line, " ")
	targetStart, _ := strconv.Atoi(values[0])
	sourceStart, _ := strconv.Atoi(values[1])
	rng, _ := strconv.Atoi(values[2])
	return &Mapping{
		TargetStart: targetStart,
		SourceStart: sourceStart,
		Range:       rng,
	}
}

func ReadSeeds(line string) []int {
	var seeds []int
	line = strings.TrimPrefix(line, "seeds: ")
	for _, seed := range strings.Split(line, " ") {
		seedNum, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedNum)
	}
	return seeds
}
