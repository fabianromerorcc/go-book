// Getting together some pointers snippets that are on chapter 2
package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1" (the value of x)
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	var y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	p = f()
	fmt.Println(f() == f()) // "false"

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
