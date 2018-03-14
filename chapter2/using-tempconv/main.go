// This is for making use of tempconv package
package main

import (
	"fmt"
	"github.com/fabianromerorcc/go-book/chapter2/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

	fmt.Printf("Kelvin boiling temperature: %v\n", tempconv.CToK(tempconv.BoilingC))
}
