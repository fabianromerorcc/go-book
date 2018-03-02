//Here are the echo exercises from 1.1 to 1.3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Exercise 1.1:")
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println()
	fmt.Println("Exercise 1.2:")

	for i, arg := range os.Args[1:] {
		fmt.Printf("%d - %s\n", i, arg)
	}

	fmt.Println()
	fmt.Println("Exercise 1.3")
	//Time for echo1
	start := time.Now()
	s = ""
	sep = ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time for echo1: %d\n", time.Since(start).Nanoseconds())

	fmt.Println()
	//Echo2
	start = time.Now()
	s = ""
	sep = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time for echo2: %d\n", time.Since(start).Nanoseconds())

	fmt.Println()
	//Echo3
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
	fmt.Printf("Time for echo3: %d\n", time.Since(start).Nanoseconds())

	fmt.Println()
	//Echo3 join directly in Println
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Time for echo3 (join inside Println): %d\n", time.Since(start).Nanoseconds())

	fmt.Println()
	//Echo3 println
	start = time.Now()
	fmt.Println(os.Args[1:])
	fmt.Printf("Time for echo3 (only Println of Args slice): %d\n", time.Since(start).Nanoseconds())
}
