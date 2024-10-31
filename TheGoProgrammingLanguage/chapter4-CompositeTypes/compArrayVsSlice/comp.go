package main

import (
	"fmt"
	"time"
)

const LONG = 1000000000

func main() {
	// Poblar un array de 10000 elementos
	startArray := time.Now()
	var arr [LONG]int
	for i := 0; i < LONG; i++ {
		arr[i] = i
	}
	durationArray := time.Since(startArray)
	fmt.Println("Tiempo para poblar el array:", durationArray)

	// Poblar un slice de 10000 elementos
	startSlice := time.Now()
	slice := make([]int, LONG)
	for i := 0; i < LONG; i++ {
		slice[i] = i
	}
	durationSlice := time.Since(startSlice)
	fmt.Println("Tiempo para poblar el slice:", durationSlice)
}
