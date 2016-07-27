package codility

import (
	"testing"
	"../"
)

func TestBinaryGap(t *testing.T)  {

	tests := map[int]int {
		9 : 2,
		15 : 0,
		20 : 1,
		529 : 4,
	}

	oBinaryGap := codility.NewBinaryGap()

	for index, expected := range tests {
		got := oBinaryGap.Solution(index)

		if(expected != got) {
			t.Error(
				"For", index,
				"expected", expected,
				"got", got,
			)
		}
	}
}
