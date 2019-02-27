//Ck converts its numeric argument to Celsius and Kelvins.
package main

import (
	"./tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ck:%v\n", err)
			os.Exit(1)
		}

		k := tempconv.Kelvins(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s \n",
			k, tempconv.KToC(k), c, tempconv.CToK(c))
	}
}
