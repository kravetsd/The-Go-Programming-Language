package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	echo3()
}

func echo3() {
	fl, err := os.Open(".ch1/some.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "echo2: %v\n", err)
	}
	println(fl.Name())
	text := bufio.NewScanner(fl)
	for text.Scan() {
		fmt.Println(text.Text())
	}
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(strings.Join(os.Args[1:], " "), "this is from echo3 function")
}
