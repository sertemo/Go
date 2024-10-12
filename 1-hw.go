package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func divmod(a int, b int) (int, int) {
	// Retorna el resultado de la división y el resto
	return a / b, a % b
}

func incrementar(a *int) {
	fmt.Println("Incrementando:", *a)
	*a = *a + 1
	fmt.Println("Incrementado:", *a)
}

func main() {
	// Comentarios de linea
	fmt.Println("Hello, World!")

	var diferente bool = (5 != 6)
	fmt.Println(diferente)

	// Loop
	var n int = 0
	for i := 0; i < 10_000_000; i++ {
		n += 1
	}
	//fmt.Println(n)

	// Usamos la función multiply
	var resultado int
	resultado = multiply(5, 6)
	fmt.Println(resultado)

	// Usamos la función divmod
	div, mod := divmod(45, 6)
	fmt.Println(div, mod)

	// Definimos un slice
	var nombres []string
	nombres = []string{"juan", "marcos", "pedro"}
	fmt.Println(nombres[0])

	nombres = append(nombres, "jose")

	// Usamos la función incrementar de punteros
	num := 5
	incrementar(&num)
	fmt.Println("Incrementado:", num)

}
