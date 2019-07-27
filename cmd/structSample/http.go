package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HttpTest() {
	client := &http.Client{}

	req, err := http.NewRequest("post", "http://www.baidu.com", strings.NewReader("key=value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "MyClient")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}


func PrintHello() {
	fmt.Println("hello world")
}
