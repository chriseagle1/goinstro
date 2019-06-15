package main

import (
	"bufio"
	"fmt"
	"gointro/cmd/pipeline"
	"os"
)

func main() {
	const filename = "small.in"
	const count = 64

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()   //程序结束时关闭文件

	p := pipeline.RandomSource(count)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	r := pipeline.ReadSource(file, -1)

	for v := range r {
		fmt.Println(v)
	}

}

func mergedemo() {
	p := pipeline.Merge(
		pipeline.Sort(pipeline.ArraySource(2,3,5,8,1,7,4)),
		pipeline.Sort(pipeline.ArraySource(0,9,7,3,2,5,4)))

	for v := range p{
		fmt.Println(v)
	}
}
