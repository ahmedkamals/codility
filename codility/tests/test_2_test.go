package codility

import (
	"testing"
	"codility/codility"
	"reflect"
)

type Test2Data struct {
	k int
	data []int
	expected int
}

func TestTest2(t *testing.T) {
	oTest2 := codility.NewTest2()

	tests := map[int]*Test2Data{
		1: &Test2Data{
			k: 1,
			data: []int{3, 5, 6, 3, 3, 5},
			expected: 4,
		},
		2: &Test2Data{
			k: 1,
			data: []int{3, 5, 6, 3, 2, 5},
			expected: 2,
		},
	}

	for index, testCase := range tests {

		got := oTest2.Solution(testCase.data)

		if(!reflect.DeepEqual(testCase.expected, got)) {
			t.Error(
				"For", index,
				"expected", testCase.expected,
				"got", got,
			)
		}
	}
}
