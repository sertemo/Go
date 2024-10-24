package main

import (
	"fmt"
	"os"
	"strings"
	//"chapter-1/dup1"
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

	var persona1 Persona

	fmt.Printf("Persona: %v\n", persona1)

	persona2 := Persona{
		Nombre: "Juan",
		Edad:   30,
	}

	fmt.Println(persona2)

	p := new(Persona)

	fmt.Println(p)

	// Creamos un mapa

	var miMapa2 map[string]int = make(map[string]int)

	miMapa2["uno"] = 1
	fmt.Println(miMapa2)
}
