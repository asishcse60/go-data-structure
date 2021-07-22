package PriorityQueue
import "errors"

type PriorityQueue struct {
	High []int
	Low []int
}

func (q *PriorityQueue) Enqueue(elem int, highPriority  bool)  {
	if highPriority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

func (q *PriorityQueue) Dequeue()(int, error)  {
	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}

	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("empty queue")
}

func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}
func (q *PriorityQueue) IsEmpty() bool{
	return q.Length() == 0
}
func (q *PriorityQueue) First() (int, error) {
	return q.Peek()
}
func (q *PriorityQueue) Last() (int, error) {
	if q.IsEmpty(){
		return 0, errors.New("queue empty")
	}
	if len(q.High) != 0 {
		return q.High[len(q.High)-1], nil
	}
	if len(q.Low) != 0 {
		return q.Low[len(q.Low)-1], nil
	}
	return 0, nil
}
func (q *PriorityQueue) Peek() (int, error){
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}
	return 0, errors.New("empty queue")
}