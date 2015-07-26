// glc2-1.go
package main

import (
	"sort"
)

func getTheBestEngineer(names []string, votes []string) string {
	var winner string

	voteCount := make(map[string]int)
	var maxVotes int

	for index, name := range votes {
		if names[index] != name {
			voteCount[name] = voteCount[name] + 1
			if voteCount[name] > maxVotes {
				winner = name
				maxVotes = voteCount[name]
			} else if voteCount[name] == maxVotes {
				winner = ""
			}
		}
	}

	return winner
}

func minFloors(M int, heights []int) int {
	sort.Ints(heights)
	var previousHeight = heights[0]
	var successiveHeights = 0
	var solutions = make([]int, 0)

	//Loop through to see if we even need to add any floors
	for _, currentHeight := range heights {
		if currentHeight == previousHeight {
			successiveHeights++
		}

		if successiveHeights == M {
			return 0
		}
	}

	//Loop through to calculate how many floors must be added to get to the desired
	// height.
	for index, currentHeight := range heights {

		//for each short floor, add the difference between heights
		// until they are the same
		var floorsAdded int
		successiveHeights = 1
		for x := index - 1; x >= 0; x-- {
			floorsAdded += currentHeight - heights[x]

			//Increment the number of heights that are the same
			successiveHeights++

			//If we hit the target leave the loop
			if successiveHeights == M {
				x = -1
			}

		}

		if successiveHeights == M {
			solutions = append(solutions, floorsAdded)
		}
	}

	sort.Ints(solutions)
	//return the smallest number found
	return solutions[0]
}
