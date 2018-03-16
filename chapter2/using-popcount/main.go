//package popcount
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	popcountexercises "github.com/fabianromerorcc/go-book/chapter2/popcount-exercises"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, err := strconv.ParseInt(input.Text(), 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("PopCount of %d is: %d\n", n, popcountexercises.PopCount(uint64(n)))
		fmt.Printf("PopCountLoop of %d is: %d\n", n, popcountexercises.PopCountLoop(uint64(n)))
		fmt.Printf("CountShift of %d is: %d\n", n, popcountexercises.PopCountShift(uint64(n)))
		fmt.Printf("ClearRight of %d is: %d\n", n, popcountexercises.PopCountClearRight(uint64(n)))
	}
}
