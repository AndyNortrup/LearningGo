// glc1_test.go
package main

import (
	"testing"
)

func TestMinimumCost(t *testing.T) {
	cost := minimumCost([]int{30, 20, 20, 10})
	if cost != 0 {
		t.Fatal("Did not caluclate cost for 30, 20, 20,10")
	}

	cost = minimumCost([]int{5, 7, 3})
	if cost != 2 {
		t.Fatalf("Did not calculate cost for {5, 7, 3} expecting 2 recieved %s", cost)
	}

	cost = minimumCost([]int{6, 8, 5, 4, 7, 4, 2, 3, 1})
	if cost != 6 {
		t.Fatalf("Did not calculate cost for {6, 8, 5, 4, 7, 4, 2, 3, 1} expecting 6 recieved %s", cost)
	}

	cost = minimumCost([]int{749, 560, 921, 166, 757, 818, 228, 584, 366, 88})
	if cost != 2284 {
		t.Fatalf("Did not calculate cost for {749, 560, 921, 166, 757, 818, 228, 584, 366, 88} expecting 2284 recieved %v", cost)
	}

	altitude := []int{712, 745, 230, 200, 648, 440, 115, 913, 627, 621, 186, 222, 741, 954, 581, 193, 266, 320, 798, 745}
	cost = minimumCost(altitude)
	if cost != 6393 {
		t.Fatalf("Did not calculate cost for %v expecting 2284 recieved %v", altitude, cost)
	}
}
