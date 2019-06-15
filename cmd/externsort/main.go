package main

import (
	"bufio"
	"fmt"
	"gointro/cmd/pipeline"
	"os"
	"strconv"
)

func main() {
	p := CreateNetworkPipeline("small.in", 4)
	writeToFile("small.out", p)

	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	source := pipeline.ReadSource(bufio.NewReader(file), -1)

	printCount := 0
	for p := range source {
		fmt.Println(p)
		printCount++
		if printCount >= 100 {
			break;
		}
	}
}

func writeToFile(filename string, in <-chan int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, in)
	writer.Flush()
}

func CreatePipeline(filename string, chunkCount int) <-chan int{
	pipeline.Init()
	sortResult := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		fileSize := getFileSize(file)
		chunkSize := fileSize/chunkCount

		file.Seek(int64(i * chunkSize), 0)
		source := pipeline.ReadSource(bufio.NewReader(file), chunkSize)

		sortResult = append(sortResult, pipeline.Sort(source))
	}

	return pipeline.MergeN(sortResult...)
}

/**
	获取文件大小
 */
func getFileSize(file *os.File) int {
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := fi.Size()

	return int(fileSize)
}

func CreateNetworkPipeline(filename string, chunkCount int) <-chan int{
	pipeline.Init()
	addrResult := []string{}
	//
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		fileSize := getFileSize(file)
		chunkSize := fileSize/chunkCount

		file.Seek(int64(i * chunkSize), 0)
		source := pipeline.ReadSource(bufio.NewReader(file), chunkSize)

		sort := pipeline.Sort(source)

		addr := ":" + strconv.Itoa(7000 + i)
		pipeline.NetworkSink(addr, sort)
		addrResult = append(addrResult, addr)
	}

	sortResult := []<-chan int{}

	for _, port := range addrResult{
		p := pipeline.NetworkReadSource(port)
		sortResult = append(sortResult, p)
	}

	return pipeline.MergeN(sortResult...)
}
