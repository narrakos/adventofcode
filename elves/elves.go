package elves

import (
	utils "adventofcode"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	CarriedCalories int
}

func GetElves() []*Elf {
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

func parseCalories(input string) []*Elf {
	var elves []*Elf
	lines := strings.Split(input, "\n")

	currentElf := &Elf{0}
	elves = append(elves, currentElf)
	for _, line := range lines {
		if line == "" {
			currentElf = &Elf{}
			elves = append(elves, currentElf)
			continue
		}

		foodCalories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentElf.CarriedCalories += foodCalories
	}
	return elves
}
