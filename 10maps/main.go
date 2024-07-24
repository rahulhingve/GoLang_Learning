package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps from go lang")
	language := make(map[string]string)
	language["js"] = "javascript"
	language["tp"] = "typescript"
	language["py"] = "python"

	fmt.Println("list of alll language :", language)
	fmt.Println("for js its  :", language["js"])
	delete(language, "py")
	fmt.Println("list of alll language :", language)

	// loops//

	for key, value := range language {
		fmt.Printf("for key %v , value is %v \n", key, value)
	}

	for _, value := range language {
		fmt.Printf("for key v , value is %v \n", value)
	}
}
