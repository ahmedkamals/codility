package codility

import (
	"testing"
	"codility/codility"
	"reflect"
)

type Test1Data struct {
	k int
	data []int
	expected int
}

func TestTest1(t *testing.T) {
	oTest1 := codility.NewTest1()

	tests := map[int]*Test1Data{
		1: &Test1Data{
			k: 1,
			data: []int{1, 4, -1, 3, 2},
			expected: 4,
		},
		2: &Test1Data{
			k: 1,
			data: []int{1, -1, 3, 2},
			expected: 2,
		},
	}

	for index, testCase := range tests {

		got := oTest1.Solution(testCase.data)

		if(!reflect.DeepEqual(testCase.expected, got)) {
			t.Error(
				"For", index,
				"expected", testCase.expected,
				"got", got,
			)
		}
	}
}
