package main

import (
	"./manual"
	"fmt"
	"sync"
)

func main() {
	man := manual.ReadManual()
	fmt.Println(Part1(man))
	fmt.Println(Part2(man))
}

// Part1 Find the lowest location value by passing all seeds through the given mappings
func Part1(manual *manual.Manual) int {
	lowestLocation := -1
	for _, seed := range manual.Seeds {
		seedLocation := getSeedLocation(seed, manual.Mappings)
		if lowestLocation == -1 {
			lowestLocation = seedLocation
		} else {
			if seedLocation < lowestLocation {
				lowestLocation = seedLocation
			}
		}
	}
	return lowestLocation
}

func getSeedLocation(seed int, manual [][]*manual.Mapping) int {
	target := seed
	for _, mappings := range manual {
		source := target
		for _, mapping := range mappings {
			if source >= mapping.SourceStart && source < mapping.SourceStart+mapping.Range {
				target = mapping.TargetStart + source - mapping.SourceStart
				break
			}
		}
	}
	return target
}

// Part2 Do the same as Part1 except the input seeds are ranges
func Part2(manual *manual.Manual) int {
	wg := &sync.WaitGroup{}
	lowestLocations := make(chan int, 10)
	for {
		seedStart := manual.Seeds[0]
		seedRange := manual.Seeds[1]
		wg.Add(1)
		go getLowestLocation(lowestLocations, wg, manual, seedStart, seedRange)
		if len(manual.Seeds) == 2 {
			break
		}
		manual.Seeds = manual.Seeds[2:]
	}
	wg.Wait()
	close(lowestLocations)
	lowestLocation := <-lowestLocations
	for island := range lowestLocations {
		if island < lowestLocation {
			lowestLocation = island
		}
	}
	return lowestLocation
}

func getLowestLocation(lowestLocations chan int, wg *sync.WaitGroup, manual *manual.Manual, seedStart int, seedRange int) {
	lowestLocation := -1
	for seed := seedStart; seed < seedStart+seedRange; seed++ {
		seedLocation := getSeedLocation(seed, manual.Mappings)
		if lowestLocation == -1 {
			lowestLocation = seedLocation
		} else {
			if seedLocation < lowestLocation {
				lowestLocation = seedLocation
			}
		}
	}
	lowestLocations <- lowestLocation
	wg.Done()
}
