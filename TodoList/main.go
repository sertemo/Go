package main

import (
	"fmt"
	"net/http"
	"strings"
)

var taskItems []string

func main() {
	fmt.Println("##### Welcome to our TodoList app! #####")

	http.HandleFunc("/", greetUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task/", addTask)

	http.ListenAndServe(":8080", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	if len(taskItems) == 0 {
		fmt.Fprintln(writer, "No hay tareas en la lista.")
		return
	}
	for i, task := range taskItems {
		fmt.Fprintf(writer, "%d: %s\n", i+1, task)
	}
}

func greetUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello, user and Welcome. Enjoy the app"

	fmt.Fprintln(writer, greeting)
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(writer, "Método no permitido. Por favor, usa una solicitud POST para agregar una tarea.", request.Method)
		return
	}

	// Parsear el formulario
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Error al procesar el formulario.")
		return
	}

	// Obtener la nueva tarea del formulario
	newTask := request.FormValue("task")
	newTask = strings.TrimSpace(newTask)
	if newTask == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "La tarea no puede estar vacía.")
		return
	}

	// Agregar la nueva tarea a la lista
	taskItems = append(taskItems, newTask)
	fmt.Fprintf(writer, "Tarea agregada: %s\n", newTask)
}
