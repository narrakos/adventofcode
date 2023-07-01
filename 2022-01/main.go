package main

import (
	"os"
	"strconv"
	"strings"
)

type elf struct {
	carriedCalories int
}

func (e *elf) addToCarriedCalories(calories int) {
	e.carriedCalories += calories
}

func main() {
	input := readInputFile()
	elves := parseCalories(input)
	elfWithMostCalories := findElfWithMostCalories(elves)

	println(elfWithMostCalories.carriedCalories)
}

func findElfWithMostCalories(elves []*elf) *elf {
	elfWithMostCalories := elves[0]

	for _, elf := range elves {
		if elfWithMostCalories.carriedCalories < elf.carriedCalories {
			elfWithMostCalories = elf
		}
	}

	return elfWithMostCalories
}

func parseCalories(input string) []*elf {
	var elves []*elf
	lines := strings.Split(input, "\n")

	currentElf := &elf{0}
	elves = append(elves, currentElf)
	for _, line := range lines {
		if line == "" {
			currentElf = &elf{}
			elves = append(elves, currentElf)
			continue
		}

		foodCalories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentElf.carriedCalories += foodCalories
	}
	return elves
}

func readInputFile() string {
	data, err := os.ReadFile("./2022-01/input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
