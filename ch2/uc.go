package main

import (
	"./unitconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unit conv %v\n", err)
			os.Exit(1)
		}

		f := unitconv.Feets(t)
		m := unitconv.Metres(t)
		fmt.Printf("%s = %s ; %s = %s\n", f, unitconv.FToM(f), m, unitconv.MToF(m))

		k := unitconv.Kilograms(t)
		p := unitconv.Pounds(t)
		fmt.Printf("%s = %s ; %s = %s\n", k, unitconv.KToP(k), p, unitconv.PToK(p))
	}
}
