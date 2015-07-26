// glc2-1_test.go
package main

import (
	"testing"
)

func TestGetTheBestEngineer(t *testing.T) {
	engineers := []string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes := []string{"Jun", "Chihiro", "Toshi", "Toshi"}

	winner := getTheBestEngineer(engineers, votes)
	expected := "Toshi"
	if winner != expected {
		t.Fatalf("Failed test 0. Expected %v Receieved %v", expected, winner)
	}

	engineers = []string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes = []string{"Teru", "Teru", "Jun", "Jun"}

	winner = getTheBestEngineer(engineers, votes)
	expected = ""
	if winner != expected {
		t.Fatalf("Failed test 1. Expected '%v' Receieved '%v'", expected, winner)
	}

	engineers = []string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes = []string{"Toshi", "Toshi", "Jun", "Jun"}

	winner = getTheBestEngineer(engineers, votes)
	expected = "Jun"
	if winner != expected {
		t.Fatalf("Failed test 2. Expected '%v' Receieved '%v'", expected, winner)
	}

	engineers = []string{"Toshi", "Jun"}
	votes = []string{"Toshi", "Jun"}

	winner = getTheBestEngineer(engineers, votes)
	expected = ""
	if winner != expected {
		t.Fatalf("Failed test 3. Expected '%v' Receieved '%v'", expected, winner)
	}

	engineers = []string{"PhpLove", "phplove", "phpLove", "Phplove"}
	votes = []string{"phpLove", "phpLove", "phpLove", "PhpLove"}

	winner = getTheBestEngineer(engineers, votes)
	expected = "phpLove"
	if winner != expected {
		t.Fatalf("Failed test 4. Expected '%v' Receieved '%v'", expected, winner)
	}

}

func TestMinFloors(t *testing.T) {
	var numJumps int
	var floors []int

	numJumps = 2
	floors = []int{1, 2, 1, 4, 3}

	result := minFloors(numJumps, floors)
	expected := 0

	if result != expected {
		t.Fatalf("Failed Test 0 - Expected: %v Recieved: %v", expected, result)
	} else {
		t.Log("Passed Test 0\n")
	}

	numJumps = 3
	floors = []int{1, 3, 5, 2, 1}
	expected = 2
	result = minFloors(numJumps, floors)

	if result != expected {
		t.Fatalf("Failed Test 1 - Expected: %v Recieved: %v", expected, result)
	} else {
		t.Log("Passed Test 1\n")
	}

	numJumps = 1
	floors = []int{43, 19, 15}
	expected = 0
	result = minFloors(numJumps, floors)

	if result != expected {
		t.Fatalf("Failed Test 2 - Expected: %v Recieved: %v", expected, result)
	} else {
		t.Log("Passed Test 2\n")
	}

	numJumps = 12
	floors = []int{25, 18, 38, 1, 42, 41, 14, 16, 19, 46, 42, 39, 38, 31, 43, 37, 26, 41, 33, 37, 45, 27, 19, 24, 33, 11, 22, 20, 36, 4, 4}
	expected = 47
	result = minFloors(numJumps, floors)

	if result != expected {
		t.Fatalf("Failed Test 2 - Expected: %v Recieved: %v", expected, result)
	} else {
		t.Log("Passed Test 2\n")
	}
}
