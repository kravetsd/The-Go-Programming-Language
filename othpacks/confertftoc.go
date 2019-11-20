package othpacks

import "fmt"

func ConvertFtoC() {
	const boilingF = 212.0
	const freezF = 32.0
	fmt.Printf("Boiling temperature %g = %g", boilingF, Ftoc(boilingF))
	fmt.Printf("Freezing temperature %g=%g", freezF, Ftoc(freezF))
}
