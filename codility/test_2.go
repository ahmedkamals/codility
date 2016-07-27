package codility

type Test2 struct {
}

func NewTest2() *Test2 {
return new(Test2)
}

func (self *Test2) Solution(A []int) int {

	pairsCount := 0

	var subIndex int
	length := len(A)

	for index, value := range A {

		subIndex = index + 1

		for subIndex < length {
			subValue := A[subIndex]
			if (value == subValue) {
				pairsCount++
			}
			subIndex++
		}
	}

	return pairsCount
}
