// Cf converts its numeric argument celsius to Fahrenheit.

package main

import (
	"chapter2/popcount"
	"chapter2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

	fmt.Println(popcount.PopCount(42))

	x := "hello"
	for _, x := range x {
		fmt.Printf("%v\n", x)
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x)
	}

}
