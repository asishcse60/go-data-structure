package main

import (
	"fmt"
	"github.com/asishcse60/go-data-structure/Queue"
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

	fmt.Println("Queue")
	queue := Queue.Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("Print queue")
	fmt.Println(queue.Elements)
	v, _:=queue.Peek()
	fmt.Printf("Peek: %v\n", v)
	v , _= queue.Dequeue()
	fmt.Printf("Dequeue: %v\n", v)
	v , _= queue.First()
	fmt.Printf("First: %v\n", v)
	v , _= queue.Last()
	fmt.Printf("Last: %v\n", v)
	v = queue.Length()
	fmt.Printf("Len: %v\n", v)
	fmt.Printf("Empty: %v\n", queue.IsEmpty())

}