package main

import (
	"fmt"
	"net/http"
	"strings"
	"top-coder/glc3/dice"
)

func main() {
	http.HandleFunc("/mf", diceService)
	http.ListenAndServe(":8080", nil)
}

func diceService(w http.ResponseWriter, r *http.Request) {

	rolls := r.URL.Query()["r"]
	faces := dice.MinimumFaces(strings.Split(rolls[0], ","))
	fmt.Fprintf(w, "%d", faces)
}
