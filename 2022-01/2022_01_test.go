package main

import (
	"testing"
)

func TestCorrectAnswerForPartOne(t *testing.T) {
	correctAnswer := 71300

	calories := getCaloriesCarriedByElves()
	sortCaloriesDescending(calories)

	if calories[0] != correctAnswer {
		t.Log("Test failed! expected: ", correctAnswer, " actual: ", calories[0])
		t.Fail()
	}
}
