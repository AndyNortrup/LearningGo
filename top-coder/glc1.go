// glc1.go
package main

import ()

func minimumCost(altitudes []int) int {

	var cost int

	for i := 1; i < len(altitudes); i++ {
		if altitudes[i] > altitudes[i-1] {
			cost += altitudes[i] - altitudes[i-1]
			altitudes[i] = altitudes[i-1]
		}
	}
	return cost
}
