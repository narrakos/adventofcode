package main

import (
	"adventofcode/elves"
)

func main() {
	elvesSlice := elves.GetElves()
	elfWithMostCalories := findElfWithMostCalories(elvesSlice)
	println(elfWithMostCalories.CarriedCalories)
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
