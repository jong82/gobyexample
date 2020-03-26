package main

import "fmt"

func add ( a int, b int) int {
	return a + b
}


func main() {
	
	for i:=1; i<10; i++ {
		results:=add (i,i+1)
		fmt.Println(i,"+",i+1,"=",results)
	}
}