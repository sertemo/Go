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

	// Slices
	// Los slices no son comparables, hay que compararlos "a mano"

	s1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s1)
	s2 := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(equal(s1, s2))

	// para ver si un slice está vacío compararlo con 0
	var sl1 []int
	fmt.Println(len(sl1), sl1 == nil, sl1)

	sl2 := []int{}
	fmt.Println(len(sl2), sl2 == nil, sl2)

	// funcion make para crear un slice
	sl3 := make([]int, 5)
	fmt.Println(sl3)

	// append function
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var xx []int
	xx = append(xx, 2, 3, 5, 7, 11, 13)
	xx = append(xx, xx...)  // apendea el propio slice a si mismo
	fmt.Println("xx=", xx)

	miString := []string{"a", "", "b", "", "c"}
	fmt.Println(nonempty(miString))

	// popear un slice
	xx = xx[:len(xx)-1]
	fmt.Println("xx=", xx)

	lp := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(lp, 2))
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}