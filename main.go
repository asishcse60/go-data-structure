package main

import (
	"fmt"
	"github.com/asishcse60/go-data-structure/BinaryTree"
)

func main() {
	//fmt.Println("Array")
	//ao.ArrayOperation()
	//var myArray [8] int
	//fmt.Println(myArray)

	//fmt.Println("Sets")
	//mySet := SetPackage.NewSet()
	//mySet.Add("Earth")
	//mySet.Add("Venus")
	//mySet.Add("Mars")
	//mySet.List()
	//fmt.Println("Ok! I am print here")
	//mySet.Add("Earth")
	//mySet.List()
	//fmt.Println("Ok! I am print here")
	//fmt.Println(mySet.Contains("Mars"))
	//fmt.Println("Ok! I am print here")
	//mySet.Delete("Venus")
	//mySet.List()

	//fmt.Println("Queue")
	//queue := Queue.Queue{}
	//queue.Enqueue(1)
	//queue.Enqueue(2)
	//queue.Enqueue(3)
	//fmt.Println("Print queue")
	//fmt.Println(queue.Elements)
	//v, _:=queue.Peek()
	//fmt.Printf("Peek: %v\n", v)
	//v , _= queue.Dequeue()
	//fmt.Printf("Dequeue: %v\n", v)
	//v , _= queue.First()
	//fmt.Printf("First: %v\n", v)
	//v , _= queue.Last()
	//fmt.Printf("Last: %v\n", v)
	//v = queue.Length()
	//fmt.Printf("Len: %v\n", v)
	//fmt.Printf("Empty: %v\n", queue.IsEmpty())

	//fmt.Println("Queues Section")

	//pq := PriorityQueue.PriorityQueue{}

	//fmt.Println(pq)
	//pq.Enqueue(1, true)
	//fmt.Println(pq)
	//pq.Enqueue(10, false)
	//fmt.Println(pq)

	//elem, _ := pq.Dequeue()
	//fmt.Println(elem)
	//fmt.Println(pq)

	//pq.Enqueue(2, true)
	//fmt.Println(pq)

	//elem, _ = pq.Dequeue()
	//fmt.Println(elem)
	//fmt.Println(pq)

	//elem, _ = pq.Dequeue()
	//fmt.Println(elem)
	//fmt.Println(pq)

	//fmt.Println("Stack Section")
	//stack := Stack.Stack{}
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Push(4)
	//stack.Push(5)
	//stack.Push(6)
	//v,_:= stack.Pop()
	//fmt.Printf("pop 1: %v\n",v)
	//v,_= stack.Pop()
	//fmt.Printf("pop 2: %v\n",v)
	//stack.Push(9)
	//stack.Push(10)

	//v,_= stack.Peek()
	//fmt.Printf("peek 1: %v\n",v)
    //fmt.Printf("len: %v\n", stack.Length())

	//v,_= stack.Pop()
	//fmt.Printf("pop 3: %v\n",v)

	//fmt.Println("LinkedList Section")
	//var ll LinkedList.LinkedList
	//ll.Insert("one")
	//ll.Insert("two")
	//ll.Insert("three")
	//ll.Insert("four")
	//ll.Insert("five")
	//ll.Insert("six")
	//ll.List()
	//node := ll.Reverse()
	//ll.ListWithValue(node)
   // fmt.Println("Ok! I am printing here")

	fmt.Println("Binary Tree Section")
	myArray :=[]int{1, 2, 3, 4, 5, 6, 6, 6, 6}
	root:=&BinaryTree.Tree{}
	root = root.BuildTree(0, myArray, root)
    root.Print()

}