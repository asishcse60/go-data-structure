package Queue

import "errors"

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(elem int)  {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue()(int, error)  {
	if len(q.Elements) == 0{
		return 0, errors.New("queue empty")
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

func (q *Queue) Length() int {
	return len(q.Elements)
}
func (q *Queue) IsEmpty() bool{
	return q.Length() == 0
}
func (q *Queue) First() (int, error) {
	return q.Peek()
}
func (q *Queue) Last() (int, error) {
	if q.IsEmpty(){
		return 0, errors.New("queue empty")
	}
	return q.Elements[len(q.Elements) - 1], nil
}
func (q *Queue) Peek() (int, error){
	if q.IsEmpty(){
		return 0, errors.New("queue empty")
	}
	firstElement := q.Elements[0]
	return firstElement, nil
}