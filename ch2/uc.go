package main

import (
	"./unitconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"reflect"
)

func showConv(t float64) {
	f := unitconv.Feets(t)
	m := unitconv.Metres(t)
	fmt.Printf("%s = %s ; %s = %s\n", f, unitconv.FToM(f), m, unitconv.MToF(m))

	k := unitconv.Kilograms(t)
	p := unitconv.Pounds(t)
	fmt.Printf("%s = %s ; %s = %s\n", k, unitconv.KToP(k), p, unitconv.PToK(p))
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				//fmt.Println(reflect.TypeOf(err))
				fmt.Fprintf(os.Stderr, "unit conv %v\n", err)
				os.Exit(1)
			}
			showConv(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unit conv %v\n", err)
				os.Exit(1)
			}

			showConv(t)
		}
	}
}
