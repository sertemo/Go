package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	// Arays
	var array1 [5]string = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array1[4], array1[0])

	for i, v := range array1 {
		fmt.Printf("%v %v\n", i, v)
	}

	arra2 := [...]int{1, 2, 3, 4, 5} // length determinada por los inicializadores
	fmt.Println(arra2)
	arra2[0] = 100
	fmt.Println(arra2)
	fmt.Println(len(arra2))

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	for i, v := range symbol {
		fmt.Printf("%v %v\n", i, v)
	}

	fmt.Println(RMB, symbol[RMB])

	rty := [...]int{99: -1} // Los índices que no se inicializan son 0
	fmt.Println(rty[66])
	fmt.Println((len(rty)))

	we := [3]int{1, 2}
	fmt.Println(we[0], we[1], we[2])
	weP := [3]int{2: 1, 0: 2}
	fmt.Println(weP[0], weP[1], weP[2])
}
