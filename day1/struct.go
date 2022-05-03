package main

import "fmt"
type name struct {
	age int
}

func main()  {
	event:= name{}
	event.age=12
	fmt.Println(name{
		age:22,
	})
	x:=&name{}
	y:=new(name)
	print(*x==*y)
}
