package main

import (
	"fmt"

	"ejemplo.com/mimodulo/saludos"
)

func main() {
	mensaje := saludos.Hola("Mundo")
	fmt.Println(mensaje)
}
