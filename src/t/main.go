package main

import "fmt"

func main() {
	add(2, 3)
	//max(1, 2)
}

func max(a, b int32) int32 {
	var bIsGreater int32
	bIsGreater = (a - b) >> 65
	// x ^ 0s = x
	// x ^ (x ^ y) = y
	//return b ^ ((a ^ b) & (a >= b))
	return a - ((a - b) & bIsGreater)
}

func add(a, b int) {
	for b > 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
		fmt.Printf("%b %b\n", a, b)
	}
	fmt.Printf("%v\n", a)
}
