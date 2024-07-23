package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello from slice ")

	var moblist = []string{"apple", "sony", "mi"}
	fmt.Printf("type of %T \n", moblist)
	moblist = append(moblist, "vivo shit ", "realme shit")
	fmt.Println(moblist)
	moblist = append(moblist[1:4]) // values giving in append cant be inclusive like it cannot be shown
	fmt.Println(moblist)

	score := make([]int, 4)

	score[0] = 11
	score[1] = 19
	score[2] = 13
	score[3] = 14

	score = append(score, 15, 16, 17)
	fmt.Println(score)

	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	var numList = []string{"one", "two", "three", "four", "five"}
	fmt.Println(numList)
	index := 2
	numList = append(numList[:index], numList[index+1:]...)
	fmt.Println(numList)

}
