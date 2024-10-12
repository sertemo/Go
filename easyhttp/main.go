package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Crear un mapa para la respuesta
	response := make(map[string]string)
	response["message"] = "¡Hola, mundo!"

	// Establecer el encabezado Content-Type a application/json
	w.Header().Set("Content-Type", "application/json")

	// Convertir el mapa a JSON y escribir la respuesta
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/piopio", handler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
