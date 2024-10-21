package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		fileName := fmt.Sprintf("test/file_%d.txt", i)
		go writeTxt(fileName, i)
	}
	wg.Wait()

}

func writeTxt(fileName string, indice int) {
	defer wg.Done()

	contenido := fmt.Sprintf("Contenido del archivo %d", indice)

	// Crear el archivo
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Escribir el contenido en el archivo
	_, err = file.WriteString(contenido)

	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(5 * time.Second)
}
