package Stack

import "errors"

type Stack struct {
	Elements []int
}

func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

func (s * Stack) Pop() (int, error) {
	if s.IsEmpty(){
		return 0, errors.New("empty stack")
	}

	lastElementIndex := s.Length() - 1

	var lastElement int
	lastElement, s.Elements = s.Elements[lastElementIndex], s.Elements[:lastElementIndex]

	return lastElement, nil
}

func (s *Stack) Peek() (int, error){
	if s.IsEmpty(){
		return 0, errors.New("empty stack")
	}
	lastElementIndex := s.Length() - 1
	return s.Elements[lastElementIndex], nil
}

func (s *Stack) IsEmpty() bool{
	return len(s.Elements) == 0
}

func (s *Stack) Length() int{
	return len(s.Elements)
}