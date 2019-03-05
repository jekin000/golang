package main

import (
	"./popcount"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcountmain %v\n", err)
			os.Exit(1)
		}
		t1 := time.Now()
		count1 := popcount.PopCount(t)
		elapsed1 := time.Since(t1)

		t2 := time.Now()
		count2 := popcount.PopCount_Cycle(t)
		elapsed2 := time.Since(t2)

		t3 := time.Now()
		count3 := popcount.PopCount_1Bit(t)
		elapsed3 := time.Since(t3)

		t4 := time.Now()
		count4 := popcount.PopCount_Clear(t)
		elapsed4 := time.Since(t4)

		fmt.Printf("direct:%d has %d's 1bits; elapsed=%s\n", t, count1, elapsed1)
		fmt.Printf(" cycle:%d has %d's 1bits; elapsed=%s\n", t, count2, elapsed2)
		fmt.Printf("  1bit:%d has %d's 1bits; elapsed=%s\n", t, count3, elapsed3)
		fmt.Printf(" clear:%d has %d's 1bits; elapsed=%s\n", t, count4, elapsed4)

		//go run popcountmain.go 123
		//direct:8 has 1's 1bits; elapsed=113ns
		//cycle:8 has 1's 1bits; elapsed=66ns
		//1bit:8 has 1's 1bits; elapsed=146ns
		//clear:8 has 1's 1bits; elapsed=40ns
	}
}
