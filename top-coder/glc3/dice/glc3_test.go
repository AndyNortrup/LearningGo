package dice

import (
	"testing"
)

func TestMinimumFaces(t *testing.T) {
	args := []string{"137", "364", "115", "724"}
	expected := 14
	result := MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 0. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 0")
	}

	args = []string{"1112", "1111", "1211", "1111"}
	expected = 5
	result = MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 1. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 1")
	}

	args = []string{"24412", "56316", "66666", "45625"}
	expected = 30
	result = MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 2. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 2")
	}

	args = []string{"931", "821", "156", "512", "129", "358", "555"}
	expected = 19
	result = MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 3. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 3")
	}

	args = []string{"3", "7", "4", "2", "4"}
	expected = 7
	result = MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 4. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 4")
	}

	args = []string{"281868247265686571829977999522",
		"611464285871136563343229916655",
		"716739845311113736768779647392",
		"779122814312329463718383927626",
		"571573431548647653632439431183",
		"547362375338962625957869719518",
		"539263489892486347713288936885",
		"417131347396232733384379841536"}
	expected = 176
	result = MinimumFaces(args)
	if result != expected {
		t.Fatalf("Failed test 5. Expected %v Actual: %v", expected, result)
	} else {
		t.Log("Passed test 5")
	}

}
