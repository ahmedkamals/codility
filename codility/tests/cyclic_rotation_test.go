package codility

import (
	"testing"
	"codility/codility"
	"reflect"
)

type CyclicData struct {
	k int
	data, expected []int
}

func TestCyclicRotation(t *testing.T) {
	oCyclicRotation := codility.NewCyclicRotation()

	tests := map[int]*CyclicData {
			1: &CyclicData {
				k: 1,
				data: []int{1,2,3},
				expected: []int{3,1,2},
			},
			2: &CyclicData {
				k: 1,
				data: []int{1,2,3},
				expected: []int{3,1,2},
			},
			3: &CyclicData {
				k: 1,
				data: []int{1,2,3},
				expected: []int{3,1,2},
			},
		}

	for index, testCase := range tests {

		got := oCyclicRotation.Solution(testCase.data, testCase.k)

		if(!reflect.DeepEqual(testCase.expected, got)) {
			t.Error(
				"For", index,
				"expected", testCase.expected,
				"got", got,
			)
		}
	}
}
