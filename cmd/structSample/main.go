package main

import "fmt"

func PrintMsg(msg *struct{
	id int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

type MyInt int


func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return int(m) + other
}
func main() {
	PrintHello()

	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}
	fmt.Printf("%T\n", msg)
	PrintMsg(msg)

	var p MyInt

	if p.IsZero() {
		fmt.Println("the number is zero")
	}

	p = 1

	fmt.Println(p.Add(2))

	HttpTest()
}
