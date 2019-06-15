package main

import (
	"fmt"
	"sort"
)

func main() {
	//Create a slice
	arr := []int{3,6,8,9,1,2,4}

	sort.Ints(arr)

	for _, v := range arr {
		fmt.Println(v)
	}
}
