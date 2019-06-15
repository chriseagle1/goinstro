package main

import (
	"bufio"
	"fmt"
	"gointro/cmd/pipeline"
	"os"
)

func main() {
	p := CreatePipeline("large.in", 4)
	writeToFile("large.out", p)

	printFile("large.out")
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
