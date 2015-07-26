package dice

import (
	"fmt"
	"net/http"
	"sort"
	//
)

func main() {
	http.HandleFunc("/", diceService)
	http.ListenAndServe(":8080", nil)
}

func diceService(w http.ResponseWriter, r *http.Request) {
	//verify that all input is a string full of numbers
	for _, value := range r.URL.Query() {
		fmt.Fprintf(w, "%s", value)
	}
}

func MinimumFaces(rolls []string) int {

	//set of dice is a
	dice := make([]int, len(rolls[0]))

	//Loop through each roll
	for _, roll := range rolls {

		pips := make([]int, len(roll))
		for index, value := range roll {
			pips[index] = int(value - 48)
		}

		sort.Ints(dice)
		sort.Ints(pips)

		for x := 0; x < len(dice); x++ {
			if pips[x] > dice[x] {
				dice[x] = pips[x]
			}
		}
	}

	//Sum the values
	sum := 0
	for _, value := range dice {
		sum += value
	}
	return sum
}
