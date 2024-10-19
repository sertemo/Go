package main

import (
	"fmt"
	"os"
	"strings"

	//"chapter-1/dup1"
	"chapter-1/fetch"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}

	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)

	var miLista = []string{
		"uno", "dos", "tres",
	}

	fmt.Println(miLista)
	fmt.Println(strings.Join(miLista, ", "))

	// loopeando sobre un map
	var miMapa = make(map[string]int)

	miMapa["uno"] = 1
	miMapa["dos"] = 2
	miMapa["tres"] = 3

	for k, v := range miMapa {
		fmt.Println(k, v)
	}

	// Loopeando sobre un struct -> No se puede iterar
	type Persona struct {
		Nombre string
		Edad   int
	}

	persona := Persona{
		Nombre: "Juan",
		Edad:   30,
	}

	fmt.Println(persona)
	// dup1.Dup()
	// lissajous.Generate(os.Stdout)

	fetch.Fetch()
}
