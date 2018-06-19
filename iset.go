package iset

type IntegerSet interface {
	IsEmpty() bool
	Add(i int)
	Contains(i int) bool
	Size() int
}

func NewArrayIntegerSet() IntegerSet {
	return &ArrayIntegerSet{elements: make([]int, 0)}
}

type ArrayIntegerSet struct {
	elements []int
}

func (s *ArrayIntegerSet) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *ArrayIntegerSet) Add(i int) {
	if s.Contains(i) {
		return
	}

	s.elements = append(s.elements, i)
}

func (s *ArrayIntegerSet) Contains(i int) bool {
	for _, el := range s.elements {
		if i == el {
			return true
		}
	}

	return false
}

func (s *ArrayIntegerSet) Size() int {
	return len(s.elements)
}
