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
		fmt.Printf("direct:%d has %d's 1bits; elapsed=%s\n", t,count1,elapsed1)
		fmt.Printf(" cycle:%d has %d's 1bits; elapsed=%s\n", t,count2,elapsed2)
		//go run popcountmain.go 123
		//direct:123 has 6's 1bits; elapsed=176ns
 		// cycle:123 has 6's 1bits; elapsed=89ns
	}
}
