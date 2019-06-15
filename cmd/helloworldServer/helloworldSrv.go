package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<H1>Hello World, %s</H1>", request.FormValue("name"))
	})
	http.ListenAndServe(":8888", nil)
}
