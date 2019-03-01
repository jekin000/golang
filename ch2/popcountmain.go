package main

import (
	"./popcount"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcountmain %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d has %d's 1bits\n", t, popcount.PopCount(t))
	}
}
