package ArrayPackage

import "fmt"

func ArrayOperation() {
	//1-D Array declaration
	//var myArray [8] int
	myArray := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len : %v\n", len(myArray))
	for idx, value := range myArray{
		fmt.Printf("Index: %v and Value: %v\n", idx, value)
	}
	fmt.Println(myArray)
   // 2D Array/ multi dimension array declare
   //var my2DArray [8][8]int
   my2DArray := [][]int{{1, 2, 3, 4},{5, 6, 7, 8}}
   row := len(my2DArray)
   col := len(my2DArray[0])
   fmt.Printf("Total Row: %v\n", row)
   fmt.Printf("Total Column: %v\n", col)

   for r := 0; r<row; r++{
   		for c:=0; c<col; c++{
   			fmt.Printf("row no: %v col no: %v and value is %v\n", r, c, my2DArray[r][c])
		}
   }
   ArraySlice()
}

func ArraySlice() {
	s := make([]string, 5)
	s[0] = "a"
	s[1] = "b"
	s[2]= "c"
	s[3]="x"
	s[4]="y"
	fmt.Println(s)
	s=append(s,"d")
 	fmt.Println(s)
	sp1 := s[2:4]
	fmt.Printf("slice part 1: %s\n", sp1)
	sp2 := s[1:]
	fmt.Printf("slice part 1: %s\n", sp2)
	sp3 := s[:3]
	fmt.Printf("slice part 1: %s\n", sp3)
	sp4 := s[0:4]
	fmt.Printf("slice part 1: %s\n", sp4)
}