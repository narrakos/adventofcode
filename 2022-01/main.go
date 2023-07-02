package main

import (
	utils "adventofcode"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	calories := getCaloriesCarriedByElves()
	sortCaloriesDescending(calories)
	partOne(calories)
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

func partOne(calories []int) {
	println("Answer for part one:", calories[0])
}
