package BinaryTree

import "fmt"

type Tree struct {
	Left *Tree
	Right *Tree
	Value int
}

func (t *Tree) BuildTree(level int, arr [] int, root *Tree)  *Tree{
	if level<len(arr){
		var treeNode = &Tree{Value: arr[level]}
		root = treeNode
		root.Left = t.BuildTree(2*level+1, arr, root.Left)
		root.Right = t.BuildTree(2*level+2, arr,root.Right)
	}
	return root
}

func (t *Tree) Print() {
	if t == nil{
		return
	}
	fmt.Println(t.Value)
	t.Left.Print()
	t.Right.Print()
}

