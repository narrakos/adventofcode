package main

import (
	utils "adventofcode"
	"os"
	"strconv"
	"strings"
)

func main() {
	calories := getCaloriesCarriedByElves()
	partOne(calories)
}

func getCaloriesCarriedByElves() []int {
	input := readInputFile()
	return parseCalories(input)
}

func readInputFile() string {
	rootDir := utils.GetProjectRootDir()
	data, err := os.ReadFile(rootDir + "/input/calories_carried_by_elves.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseCalories(input string) []int {
	var caloriesCarriedByElves []int
	lines := strings.Split(input, "\n")

	var temp = 0
	for _, line := range lines {
		if line == "" {
			caloriesCarriedByElves = append(caloriesCarriedByElves, temp)
			temp = 0
			continue
		}

		foodCalories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		temp += foodCalories
	}
	return caloriesCarriedByElves
}

func partOne(elvesSlice []int) {
	elfWithMostCalories := findElfWithMostCalories(elvesSlice)
	println("Answer for part one:", elfWithMostCalories)
}

func findElfWithMostCalories(elves []int) int {
	elfWithMostCalories := elves[0]

	for _, elf := range elves {
		if elfWithMostCalories < elf {
			elfWithMostCalories = elf
		}
	}

	return elfWithMostCalories
}
