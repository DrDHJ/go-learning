package main

import (
	"The-Go-Programming-Language/ch02/popcount"
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"time"
)

func main() {
	count := len(os.Args)
	if count < 2 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			comparePopCount(input.Text())
		}
	} else {
		for _, s := range os.Args[1:] {
			comparePopCount(s)
		}
	}
}

func comparePopCount(s string) {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
		return
	}
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = popcount.PopCountAdd(uint64(i))
	}
	result1 := popcount.PopCountAdd(x)
	elapsed1 := time.Since(start).Nanoseconds()

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = popcount.PopcountFor(uint64(i))
	}
	result2 := popcount.PopcountFor(x)
	elapsed2 := time.Since(start).Nanoseconds()

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = popcount.PopCountNaive(uint64(i))
	}
	result3 := popcount.PopCountNaive(x)
	elapsed3 := time.Since(start).Nanoseconds()

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = popcount.PopCountClear(uint64(i))
	}
	result4 := popcount.PopCountClear(x)
	elapsed4 := time.Since(start).Nanoseconds()

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = bits.OnesCount64(uint64(i))
	}
	result5 := bits.OnesCount64(x)
	elapsed5 := time.Since(start).Nanoseconds()

	fmt.Printf("PopCountAdd(%d) = %d, time = %d ns\n", x, result1, elapsed1)
	fmt.Printf("PopcountFor(%d) = %d, time = %d ns\n", x, result2, elapsed2)
	fmt.Printf("PopCountNaive(%d) = %d, time = %d ns\n", x, result3, elapsed3)
	fmt.Printf("PopCountClear(%d) = %d, time = %d ns\n", x, result4, elapsed4)
	fmt.Printf("bits.OnesCount64(%d) = %d, time = %d ns\n", x, result5, elapsed5)
}
