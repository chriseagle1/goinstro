package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string `xml:"name"`
	Age int
}


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

	pers1 := Person{"chris", 30}

	fmt.Println(pers1)

	var data []byte
	var err  error

	if data, err = xml.Marshal(pers1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	pers2 := new(Person)

	if err = xml.Unmarshal(data, pers2);err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pers2)
}
