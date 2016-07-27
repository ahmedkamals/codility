package codility

type CyclicRotation struct {
}

func NewCyclicRotation() *CyclicRotation {
	return new(CyclicRotation)
}

func (self *CyclicRotation) Solution(A []int, k int) []int{

	length := len(A)
	copy := make([]int, length)

	for index, value := range A {
		copy[(index + k) % length] = value
	}

	return copy
}