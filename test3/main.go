package main

import(
	"fmt"
)

func main(){
	x := []int{10,20,30}
	y := append(x,40,50)
	fmt.Println(x,y)
}