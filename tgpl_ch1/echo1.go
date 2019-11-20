package main

import (
	"flag"
	"fmt"
	"strings"
	"tgpl-github/testfoldermodule"
)

var n = flag.Bool("n", false, "ommit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	testfoldermodule.Subecho()

}

func Hello() string {
	//comment here
	//There is a comment just to fix syntax of the golang
	s := "Hello wotld!"
	return s
}
