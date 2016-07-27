package codility

import (
	"testing"
	"codility/codility"
	"reflect"
)

type Test3Data struct {
	k int
	data []int
	expected bool
}

func TestTest3(t *testing.T) {
	oTest3 := codility.NewTest3()

	tests := map[int]*Test3Data{
		1: &Test3Data{
			k: 1,
			data: []int{1, 5, 3, 3, 7},
			expected: true,
		},
		2: &Test3Data{
			k: 1,
			data: []int{1, 3, 5, 3, 4},
			expected: false,
		},
		3: &Test3Data{
			k: 1,
			data: []int{1, 3, 5},
			expected: true,
		},
	}

	for index, testCase := range tests {

		got := oTest3.Solution(testCase.data)

		if(!reflect.DeepEqual(testCase.expected, got)) {
			t.Error(
				"For", index,
				"expected", testCase.expected,
				"got", got,
			)
		}
	}
}
