package main

import (
	"fmt"
	"strings"
)

func main()  {
	strList := []string {
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	funcList := []func(string) string {
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	for index, str := range strList {
		for _, proc := range funcList {
			str = proc(str)
		}
		strList[index] = str
	}

	for _, res := range strList {
		fmt.Println(res)
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}


