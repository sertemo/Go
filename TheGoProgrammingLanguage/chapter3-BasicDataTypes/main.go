package main

import (
	"bytes"
	basename "chapter3/basename1"
	comma "chapter3/comma"
	"fmt"
	"math"
	"strconv"
	"time"
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

	var path string = "C:/Usuarios/usuario/Documentos/Go/src/chapter3/basename1/basename1.go"
	s := basename.Basename(path)
	s2 := basename.Basename2(path)
	fmt.Println(s, s2)

	fmt.Println(comma.Comma("1234567"))
	fmt.Println(comma.Comma("1234"))
	fmt.Println(comma.Comma("123"))
	fmt.Println(comma.Comma("12"))
	fmt.Println(comma.Comma("1"))

	s3 := "Zabcz"
	b := []byte(s3)
	fmt.Println(s3, b)
	fmt.Println(string(b))

	intArray := []int{1, 2, 3, 4, 5}
	fmt.Println(intsToString(intArray))
	fmt.Println(intArray)

	conv, _ := strconv.Atoi("1234")
	fmt.Printf("%d, %T\n", conv, conv)

	// constantes
	const noDelay time.Duration = 0
	const timeout = 5*time.Minute + 30*time.Second
	fmt.Printf("%v, %T\n", noDelay, noDelay)
	fmt.Printf("%v, %T\n", timeout, timeout)

	const (
		a = 1
		bb
		c = 2
		d
	)
	fmt.Println(a, bb, c, d)

	// iota
	type Weekday int

	const (
		Monday Weekday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	type Flags uint
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)

	fmt.Printf("%v\n", FlagUp)
	fmt.Printf("%v\n", FlagBroadcast)
	fmt.Printf("%v\n", FlagLoopback)
	fmt.Printf("%v\n", FlagPointToPoint)

	var it int8 = 1

	fmt.Println(it, it+1, it*2, it/2, it%2, -it)
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
