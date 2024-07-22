package main

import "fmt"

func main() {
	fmt.Println("arrays here ")
	var fruitList [4]string
	fruitList[0] = "rahul"
	fruitList[1] = "rahul1"
	fruitList[3] = "rahul2"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var listofmobile = [3]string{"12 pro max ", "iqoo 12", "s24 ultra"}
	fmt.Println("mopbiles list : ", listofmobile)
	fmt.Println(listofmobile)

}
