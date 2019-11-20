package main

import (
	"fmt"

	othpacks "tgpl-github/othpacks"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	fmt.Printf("boiling point = %g°F or %g°C\n", f, othpacks.Ftoc(f))
	//Output:
	//Boiling point = 212F° or 100°C
}
