package codility

type Test1 struct {
}

func NewTest1() *Test1 {
	return new(Test1)
}

func (self *Test1) Solution(A []int) int {

	length := 0
	value := A[0]
	for value != -1 {
		length++
		value = A[value]
	}

	return length + 1
}
