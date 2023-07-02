package main

import (
	"testing"
)

// 209691
func TestCorrectAnswerForPartOne(t *testing.T) {
	correctAnswer := 71300

	calories := getCaloriesCarriedByElves()
	sortCaloriesDescending(calories)

	if calories[0] != correctAnswer {
		t.Log("Test failed! expected: ", correctAnswer, " actual: ", calories[0])
		t.Fail()
	}
}

func TestCorrectAnswerForPartTwo(t *testing.T) {
	correctAnswer := 209691

	calories := getCaloriesCarriedByElves()
	sortCaloriesDescending(calories)

	caloriesCarriedByTopThreeElves := calories[0] + calories[1] + calories[2]

	if caloriesCarriedByTopThreeElves != correctAnswer {
		t.Log("Test failed! expected: ", correctAnswer, " actual: ", caloriesCarriedByTopThreeElves)
		t.Fail()
	}
}
