package codility

import "strconv"

type BinaryGap struct{
}

func NewBinaryGap() *BinaryGap {
	return new (BinaryGap)
}

/**
 * Get the number of the longest binary gap.
 * @ example, number 9 has binary representation 1001 and contains a binary gap of length 2.
 * The number 529 has binary representation 1000010001 and contains two binary gaps:
 * one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps
 */
func (self *BinaryGap) Solution(n int) int {

	//r, _ := regexp.Compile("0*(10+1)+0*");
	//match := r.FindAllStringSubmatch(str, -1)
	str := strconv.FormatInt(int64(n), 2)

	gapLength := 0
	gapsCounter := 0
	maxGapLength := 0

	previousOne := false

	for _, value := range str {

		char := string(value)

		switch char {
			case "0":
				if(true == previousOne) {
					gapLength++
				}
			break
			case "1":
				if(true == previousOne) {
					gapsCounter++
					if(maxGapLength <= gapLength) {
						maxGapLength = gapLength
					}
					gapLength = 0
				}

				previousOne = true
			break
		}
	}

	return maxGapLength
}
