package main

import "fmt"

func main() {
	fmt.Println("this is pointer tutorial ")

	// var pointer *string    // this is pointer

	// fmt.Println("value of pointer is ", pointer)

	myNumber := 23
	var pointer = &myNumber
	fmt.Println("value is ", pointer)
	fmt.Println("value is ", *pointer+2)
	fmt.Println("value is ", myNumber)
	*pointer = *pointer + 2
	fmt.Println(myNumber)
}
