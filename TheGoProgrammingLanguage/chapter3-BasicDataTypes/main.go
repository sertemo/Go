package main

import (
	"fmt"
	"math"
)

func main() {

	var u int8 = 127
	fmt.Println(u, u+1, u*u)
	fmt.Printf("%b", u)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	medals := []string{"gold", "silver", "bronze"}

	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(i+1, medals[i])
	}

	var apples int32 = 1
	var oranges int16 = 2

	compota := int(apples) + int(oranges) // Hay que convertir al mismo tipo

	fmt.Println(compota)

	f := 3.1416
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	// formatos al imprimir
	fmt.Println("############")

	o := 0666                          // octal
	fmt.Printf("%d %[1]o %#[1]o\n", o) // [1] como adverbio le dice a Printf que use el primer argumento una y otra vez
	g := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", g)

	// Runas
	fmt.Println("############")
	ascii := 'a'
	unicode := '汉'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	// floats
	var ff float32 = 16777216 // si el float32 da true si es float64 da false
	fmt.Println(ff == ff+1)

	const Avogadro = 6.02214129e23
	fmt.Println(Avogadro)

	for s := 0; s < 10; s++ {
		fmt.Printf("x = %d ex = %8.3f\n", s, math.Exp(float64(s)))
	}

	var ty float64

	fmt.Printf("0/0 = %g\n", ty/ty)

	var nan float64 = math.NaN()
	var inf float64 = math.Inf(1)
	fmt.Println(nan == nan, nan < nan, nan > nan) // siempre es FALSE
	fmt.Println(inf == inf, inf < inf, inf > inf)

}
