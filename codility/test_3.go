package codility

type Test3 struct {
}

func NewTest3() *Test3 {
	return new(Test3)
}

func (self *Test3) Solution(A []int) bool {

	swapsCount := 0
	var subIndex, subValue, subMax int
	length := len(A)

	for index, value := range A {
		subIndex = index + 1

		if (subIndex < length) {
			subMax = A[subIndex]
		}

		for subIndex < length {
			subValue = A[subIndex]
			subIndex++

			if (subValue < value) {
				swapsCount++

				if (subMax >= subValue) {
					swapsCount = 1
				}
			}

			if (subMax < subValue) {
				subMax = subValue
			}
		}
	}

	return swapsCount <= 1
}
