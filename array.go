package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	var  arr =  [5]int {1,2,3,4,5}
	var i int
	for i=0;i<len(arr); i++{
		fmt.Println("printing element",arr[i])
	}
	var value int
	for i , value = range arr{
		fmt.Println("range",value)
	}
	for _,value=range arr{
		fmt.Println("blank range", value)
	}

}