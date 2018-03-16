//package popcount
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	fmt.Println("Computing lookup table...")
	for i := range pc {
		fmt.Printf("Value of i: %v, and i/2: %v. Value of pc[%v]: %v, and pc[%v](i/2): %v. byte(i&1): %v\n", i, i/2, i, pc[i], i/2, pc[i/2], byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("Value set on pc[%v]: %v\n", i, pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println("PC Array:")
	fmt.Println(pc)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, err := strconv.ParseInt(input.Text(), 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("PopCount of %d is: %d\n", n, PopCount(uint64(n)))
	}
}
