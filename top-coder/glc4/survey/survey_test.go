package survey

import (
	"testing"
)

func TestBigShoppers(t *testing.T) {

	shoppers := []int{5, 100, 10, 7, 5}
	survey := [][]int{{3, 3}, {97}, {9, 9, 9, 9, 9}, {1, 2, 3}, {3, 3, 3}}
	expected := []int{1, 97, 5, 0, 0}

	var fail bool = false
	for index, _ := range shoppers {
		result := CountBigShoppers(shoppers[index], survey[index])
		if expected[index] != result {

			t.Logf("Failed test %v \t Expected: %v \t Actual: %v \n",
				index,
				expected[index],
				result)
			fail = true
		} else {
			t.Logf("Passed test %v", index)
		}

		if fail {
			t.Fail()
		}
	}
}

func TestNetBigShoppers(t *testing.T) {

}
