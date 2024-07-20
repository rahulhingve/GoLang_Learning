package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input now type here  "
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter")
	// comma ok systax / / err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks ", input)
	fmt.Printf("type of %T ", input)

}
