package main

import (
	"flag"
	"fmt"
)

func style1()  {
	var method = flag.String("method", "default", "cli param for string")
	var value = flag.Int("value", -1, "cli param for int")

	flag.Parse()
	fmt.Println(*method, *value)
}

func style2() {
	var method string
	var value int

	flag.StringVar(&method, "m", "default", "cli param for string")
	flag.IntVar(&value, "v", -1, "cli param for int")

	flag.Parse()

	fmt.Println(method, value)
}


func main() {
	//style1()
	style2()
}
