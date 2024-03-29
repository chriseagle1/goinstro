package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "cli demo");

func main() {
	flag.Parse()

	var skill = map[string]func() {
		"fire": func() {
			fmt.Println("chicken fired")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[(*skillParam)]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
