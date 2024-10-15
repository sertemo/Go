package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Result struct {
	Index int
	Body  string
}

func main() {
	fileName := "urls.txt"
	urls, err := getUrlListFromFile(fileName)
	if err != nil {
		fmt.Println("Error al obtener la lista de URLs:", err)
		return
	}

	var wg sync.WaitGroup
	results := make(chan Result) // Canal sin buffer para comunicación entre productores y consumidores

	// Iniciar goroutines consumidoras para escribir archivos
	numConsumers := 5 // Número de goroutines consumidoras
	var wgConsumers sync.WaitGroup
	wgConsumers.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go func() {
			defer wgConsumers.Done()
			for result := range results {
				err := writeHTML(result.Body, result.Index)
				if err != nil {
					fmt.Printf("Error al escribir el archivo %d: %v\n", result.Index, err)
				} else {
					fmt.Printf("Archivo %d escrito correctamente.\n", result.Index)
				}
			}
		}()
	}

	// Iniciar goroutines productoras para hacer solicitudes HTTP
	for index, url := range urls {
		wg.Add(1)
		go func(url string, index int) {
			defer wg.Done()
			body := makeRequest(url)
			if body != "" {
				results <- Result{Index: index, Body: body}
			}
		}(url, index)
	}

	// Cerrar el canal de resultados una vez que todos los productores hayan terminado
	go func() {
		wg.Wait()
		close(results)
	}()

	// Esperar a que los consumidores terminen
	wgConsumers.Wait()
}

func makeRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error al realizar la solicitud a %s: %v\n", url, err)
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error al leer la respuesta de %s: %v\n", url, err)
		return ""
	}

	return string(bodyBytes)
}

func writeHTML(body string, index int) error {
	fileName := fmt.Sprintf("pagina_%d.html", index)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(body)
	if err != nil {
		return err
	}

	return nil
}

func getUrlListFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
