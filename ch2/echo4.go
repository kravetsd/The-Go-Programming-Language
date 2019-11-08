package echo

import (
	"flag"
	"fmt"
	"strings"
	"tgpl/ch2/test_folder_module"
)

var n = flag.Bool("n", false, "ommit trailing newline")
var sep = flag.String("s", " ", "separator")

func Echo() {
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
