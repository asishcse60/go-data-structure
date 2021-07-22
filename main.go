package main

import (
	"fmt"
	"github.com/asishcse60/go-data-structure/SetPackage"
)

func main() {
	//fmt.Println("Array")
	//ao.ArrayOperation()
	//var myArray [8] int
	//fmt.Println(myArray)

	fmt.Println("Sets")
	mySet := SetPackage.NewSet()
	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.List()
	fmt.Println("Ok! I am print here")
	mySet.Add("Earth")
	mySet.List()
	fmt.Println("Ok! I am print here")
	fmt.Println(mySet.Contains("Mars"))
	fmt.Println("Ok! I am print here")
	mySet.Delete("Venus")
	mySet.List()
}