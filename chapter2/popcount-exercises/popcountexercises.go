package popcountexercises

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
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

// PopCountLoop uses a loop instead of a single expression
func PopCountLoop(x uint64) int {
	var count byte
	for i := byte(0); i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

// PopCountShift counts bits by shifting its argument through 64 bit positions
func PopCountShift(x uint64) int {
	var count byte
	for x != 0 {
		count += byte(x & 1)
		x >>= 1
	}
	return int(count)
}

// PopCountClearRight uses x & (x-1) to clear the rihtmost non-zero bit of x
func PopCountClearRight(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

//TODO Create a recursive version of ClearRight and assess its performance
