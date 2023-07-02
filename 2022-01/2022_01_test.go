package main

import (
	"adventofcode/elves"
	"testing"
)

func TestCorrectAnswerForPartOne(t *testing.T) {
	correctAnswer := 71300

	elvesSlice := elves.GetElves()
	elfWithMostCalories := findElfWithMostCalories(elvesSlice)

	if elfWithMostCalories.CarriedCalories != correctAnswer {
		t.Log("Test failed! expected: ", correctAnswer, " actual: ", elfWithMostCalories.CarriedCalories)
		t.Fail()
	}
}
