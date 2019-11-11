package main

import (
	"flag"
	"fmt"
	"strings"

	test_folder_module "github.com/oth_packs"
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
	s := "Hello wotld!"
	return s

}
