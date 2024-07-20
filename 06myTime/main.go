package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time machine of go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04"))

	createDate := time.Date(2023, time.Now().Month(), 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 Monday 15:04"))

}
