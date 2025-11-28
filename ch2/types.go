package main

import (
	"fmt"
	"math"
)

func main() {
	var flag bool = true
	var byt uint8 = 0
	var min64 int64 = math.MinInt64
	var max64 int64 = math.MaxInt64

	byt = 123

	fmt.Println(flag)
	fmt.Println(byt)
	fmt.Println(1.0 / float64(byt))
	fmt.Println(min64)
	fmt.Println(max64)
}
