package main

import (
	"adventofcode/elves"
)

func main() {
	elvesSlice := elves.GetElves()
	partOne(elvesSlice)
}

func partOne(elvesSlice []*elves.Elf) {
	elfWithMostCalories := findElfWithMostCalories(elvesSlice)
	println("Answer for part one:", elfWithMostCalories.CarriedCalories)
}

func findElfWithMostCalories(elves []*elves.Elf) *elves.Elf {
	elfWithMostCalories := elves[0]

	for _, elf := range elves {
		if elfWithMostCalories.CarriedCalories < elf.CarriedCalories {
			elfWithMostCalories = elf
		}
	}

	return elfWithMostCalories
}
