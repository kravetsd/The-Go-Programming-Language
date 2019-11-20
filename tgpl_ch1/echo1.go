package main

import (
	"flag"
	"fmt"
	"strings"
	test_folder_module "tgpl-github/oth_packs"
)

var n = flag.Bool("n", false, "ommit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	test_folder_module.Subecho()

}

func Hello() string {
	//comment here
	//There is a comment just to fix syntax of the golang
	s := "Hello wotld!"
	return s

}
