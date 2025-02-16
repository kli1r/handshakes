package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var k uint64
	fmt.Fscan(os.Stdin, &k)

	if float64(k) > 0.5*(1+math.Sqrt(147573952589676412929)) {
		fmt.Println("Переполнение")
		return
	}

	fmt.Println((k * (k - 1)) / 2)
}
