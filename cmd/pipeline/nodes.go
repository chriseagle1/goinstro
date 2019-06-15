package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init()  {
	startTime = time.Now()
}

func ArraySource(arr...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range arr{
			out <- v
		}
		close(out)
	}()
	return out
}

/**
	排序
 */
func Sort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		var a []int
		for v := range in{
			a = append(a, v)
		}
		fmt.Println("Read done to memory:", time.Now().Sub(startTime))
		sort.Ints(a)
		fmt.Println("Sort done in memory:", time.Now().Sub(startTime))
		for _, el := range a{
			out <- el
		}
		close(out)
	}()
	return out
}

/**
	合并
 */
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 2048)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}

		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}


/**
	读数据
 */
func ReadSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		readSize := 0
		for {
			n, err := reader.Read(buffer)
			readSize += n
			if n > 0 {
				v := binary.BigEndian.Uint64(buffer)
				out <- int(v)
			}
			if err != nil || (chunkSize != -1 && readSize >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

/**
	写数据
 */
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in{
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

/**
	产生随机数
 */
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

/**
	多个channel merge
 */
func MergeN(inputs... <-chan int) <-chan int {
	count := len(inputs)

	if count == 1 {
		return inputs[0]
	}

	m := count/2

	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
