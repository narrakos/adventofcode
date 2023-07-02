package main

import (
	"testing"
)

func TestCorrectAnswerForPartOne(t *testing.T) {
	correctAnswer := 71300

	calories := getCaloriesCarriedByElves()
	mostCaloriesCarriedByElf := findElfWithMostCalories(calories)

	if mostCaloriesCarriedByElf != correctAnswer {
		t.Log("Test failed! expected: ", correctAnswer, " actual: ", mostCaloriesCarriedByElf)
		t.Fail()
	}
}
