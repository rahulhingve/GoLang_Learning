package main

import "fmt"

func main() {
	var username string = "rahul"
	fmt.Println(username)
	fmt.Printf("Variavble name: %T \n", username)

	var checkit bool = true // or false
	fmt.Println(checkit)
	fmt.Printf("type of : %T \n", checkit)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("type of : %T \n", smallValue)

	var smallFloat float64 = 255.454555555555555
	fmt.Println(smallFloat)
	fmt.Printf("type of : %T \n", smallFloat)

	//value

	var intvariable int
	fmt.Println(intvariable)
	fmt.Printf("type of : %T \n", intvariable)

	var strvariable string
	fmt.Println(strvariable)
	fmt.Printf("type of : %T \n", strvariable)

	//implicit type

	var nambol = "rahulhai"
	fmt.Println(nambol)

	//no declaration

	numberofuser := 30000
	fmt.Println(numberofuser)
}
