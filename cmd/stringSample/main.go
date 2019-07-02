package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world"

	fmt.Println(strings.Contains(str, "o"), strings.Contains(str, "?"))

	fmt.Println(strings.Index(str, "o"))

	str2 := "1#2#4#455#3#27"

	splitStr := strings.Split(str2, "#")
	fmt.Println(splitStr)

	joinStr := strings.Join(splitStr, "$")

	fmt.Println(joinStr)

	fmt.Println(strings.ContainsRune(str, 'e'))

	fmt.Println(strconv.Itoa(48))
	fmt.Println(strconv.Atoi("23a4"))
}
