package main

import (
	utils "adventofcode"
	"sort"
	"strconv"
	"strings"
)

func main() {
	calories := getCaloriesCarriedByElves()
	sortCaloriesDescending(calories)
	partOne(calories)
	partTwo(calories)
}

func sortCaloriesDescending(calories []int) {
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
}

func getCaloriesCarriedByElves() []int {
	input := readInputFile()
	return parseCalories(input)
}

func readInputFile() string {
	return utils.ReadInputFile("calories_carried_by_elves.txt")
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

func partOne(calories []int) {
	println("Answer for part one:", calories[0])
}

func partTwo(calories []int) {
	println("Answer for part two:", calories[0]+calories[1]+calories[2])

}
