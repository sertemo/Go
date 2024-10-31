package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	comp1 := "x"
	comp2 := "X"
	c1 := sha256.Sum256([]byte(comp1))
	c2 := sha256.Sum256([]byte(comp2))
	fmt.Printf("%v - %v\n", comp1, comp2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// Poner a 0 un array de longitud fija
	zero(&c1)
	fmt.Printf("%v\n", c1)
}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
