package LinkedList

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}
type Node struct {
	value string
	Next *Node
}

func (l *LinkedList) Insert(elem string){
	node:=Node{value: elem, Next: l.Head}
	l.Head = &node
	l.Size++
}
func (l *LinkedList) Search(elem string) *Node{
	if l.Head == nil{
		return nil
	}
	current:=l.Head
	for current != nil{
		if current.value == elem{
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) Reverse() *Node{
	if l.Head == nil{
		return nil
	}
	cur := l.Head
	var prev *Node = nil
	for cur != nil {
		var nextNode = cur.Next
		cur.Next = prev
		prev = cur
		cur = nextNode
	}
	return prev
}

func (l *LinkedList) HasCycle() bool{
	slow:=l.Head
	fast:=l.Head.Next
	if slow == nil || fast==nil{
		return false
	}
	for slow != fast {
		if fast == nil || fast.Next == nil{
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func (l *LinkedList) Delete(elem string){
	if l.Head == nil{
		return
	}
	prev:= l.Head
	cur := l.Head
	for cur!=nil {
		if cur.value == elem{
			prev.Next = cur.Next
			l.Size--
		}
		prev = cur
		cur = cur.Next
	}
}

func (l *LinkedList) DeleteFirst(){
	if l.Head == nil{
		return
	}
	l.Head = l.Head.Next
	l.Size--
}
func (l *LinkedList) List() {
	current:=l.Head
	for current!=nil{
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

func (l *LinkedList) ListWithValue(node *Node) {
	current:=node
	for current!=nil{
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}