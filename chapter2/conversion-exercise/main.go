// Unit conversion - Exercise 2.2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fabianromerorcc/go-book/chapter2/lenconv"
	"github.com/fabianromerorcc/go-book/chapter2/tempconv"
	"github.com/fabianromerorcc/go-book/chapter2/weightconv"
)

func main() {
	numbers := os.Args[1:]
	if len(numbers) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convertAndPrint(input.Text())
		}
	} else {
		for _, n := range numbers {
			convertAndPrint(n)
		}
	}
}

func convertAndPrint(arg string) {
	n, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unit-conversion: %v\n", err)
		os.Exit(1)
	}
	// Print number in every type created for conversion
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	m := lenconv.Meter(n)
	ft := lenconv.Feet(n)
	kg := weightconv.Kilogram(n)
	p := weightconv.Pound(n)
	fmt.Println("Temperature conversions")
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Println("Length conversions")
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToFt(m), ft, lenconv.FtToM(ft))
	fmt.Println("Weight conversions")
	fmt.Printf("%s = %s, %s = %s\n", kg, weightconv.KgToP(kg), p, weightconv.PToKg(p))

}
