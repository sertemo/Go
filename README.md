# Lenguaje Golang
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Repositorio para aprender conceptos del lenguaje Go siguiendo algunos tutorial de Youtube:
- [Alex Mux](https://www.youtube.com/watch?v=8uiZC0l4Ajw&list=WL&index=2)
- [freeCodeCamp.org](https://www.youtube.com/watch?v=un6ZyFkqFKo&list=WL&index=3&t=341s)


## Tutorial básico para comenzar con Go

### 1. Instalación de Go
Descargar e instalar:

- Visita la [página oficial](https://golang.org/dl/)
- Descarga el instalador adecuado para tu sistema operativo (Windows, macOS, Linux).
- Sigue las instrucciones del instalador.

Verificar la instalación:

Abre una terminal y ejecuta:

```bash
go version
```

Deberías ver algo similar a:

```bash
go version go1.23.2 windows/amd64
```

### 2. Configuración del entorno
Go utiliza dos variables de entorno importantes:

- **GOROOT**: Directorio donde está instalado Go. Normalmente se configura automáticamente.
- **GOPATH**: Directorio de trabajo donde estarán tus proyectos y dependencias.

Por defecto, GOPATH suele ser $HOME/go. Asegúrate de añadir $GOPATH/bin a tu variable PATH.

### 3. Tu primer programa en Go
Crea un directorio para tu proyecto, por ejemplo `myProject`.

```bash
mkdir myProject
cd myProject
```

Crea un archivo llamado `main.go`:

```go
Copiar código
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, mundo!")
}
```

Ejecuta el programa:

```bash
go run main.go
```

Deberías ver:

¡Hola, mundo!

### 4. Estructura básica de un programa
- package main: Indica que es un programa ejecutable.
- import: Importa paquetes necesarios.
- func main(): Punto de entrada del programa.


### 5. Variables y tipos de datos
Declaración de variables:

```go
var nombre string = "Ana"
var edad int = 25

// Declaración corta
ciudad := "Madrid"
```

Tipos de datos básicos:

- Números enteros: int, int8, int16, int32, int64
- Números de punto flotante: float32, float64
- Booleanos: bool
- Cadenas: string


### 6. Operadores y expresiones
Operadores aritméticos:

```go
suma := 10 + 5
resta := 10 - 5
producto := 10 * 5
division := 10 / 5
modulo := 10 % 3
```

Operadores relacionales:

```go
Copiar código
igual := (5 == 5)
diferente := (5 != 3)
mayor := (5 > 3)
menor := (3 < 5)
```

7. Estructuras de control
Condicional if:

```go
if edad >= 18 {
    fmt.Println("Eres mayor de edad")
} else {
    fmt.Println("Eres menor de edad")
}
```

Bucle for:

```go
for i := 0; i < 5; i++ {
    fmt.Println("Número:", i)
}
```

Bucle while (usando for):

```go
i := 0
for i < 5 {
    fmt.Println("Número:", i)
    i++
}
```

Switch:

```go
dia := "lunes"
switch dia {
case "lunes":
    fmt.Println("Inicio de semana")
case "viernes":
    fmt.Println("Fin de semana")
default:
    fmt.Println("Día intermedio")
}
```

### 8. Funciones
Definición y llamada de funciones:

```go
func saludo(nombre string) {
    fmt.Println("Hola,", nombre)
}

saludo("Carlos")
```

Funciones con retorno:

```go
func suma(a int, b int) int {
    return a + b
}

resultado := suma(3, 4)
fmt.Println("Resultado:", resultado)
```

Múltiples retornos:

```go
func operaciones(a int, b int) (int, int) {
    suma := a + b
    resta := a - b
    return suma, resta
}

s, r := operaciones(10, 5)
fmt.Println("Suma:", s, "Resta:", r)
```

### 9. Arrays y slices
Arrays (arreglos):

```go
var numeros [3]int
numeros[0] = 1
numeros[1] = 2
numeros[2] = 3
fmt.Println(numeros)
```

Slices:

```go
frutas := []string{"Manzana", "Banana", "Cereza"}
frutas = append(frutas, "Durazno")
fmt.Println(frutas)
```

### 10. Mapas (maps)
Estructura clave-valor similar a los diccionarios.

```go
edades := make(map[string]int)
edades["Ana"] = 30
edades["Luis"] = 25

fmt.Println(edades)
fmt.Println("Edad de Ana:", edades["Ana"])
```

### 11. Estructuras (structs)
Definición de tipos personalizados.

```go
type Persona struct {
    Nombre string
    Edad   int
}

p := Persona{
    Nombre: "Laura",
    Edad:   22,
}

fmt.Println(p)
fmt.Println("Nombre:", p.Nombre)
```

### 12. Punteros
Referencia a la dirección de memoria de una variable.

```go
func incrementar(a *int) {
    *a = *a + 1
}

num := 5
incrementar(&num)
fmt.Println("Incrementado:", num)
```

### 13. Concurrencia
Goroutines:

Son funciones que se ejecutan de forma concurrente.

```go
func imprimirMensaje(msg string) {
    fmt.Println(msg)
}

go imprimirMensaje("Mensaje en goroutine")
imprimirMensaje("Mensaje en main")
```

Canales:

Comunicación entre goroutines.

```go
Copiar código
messages := make(chan string)

go func() {
    messages <- "Hola desde la goroutine"
}()

msg := <-messages
fmt.Println(msg)
```

### 14. Paquetes y módulos
Crear un módulo:

```bash
go mod init ejemplo.com/mimodulo
```

Estructura de directorios:

```go
mi_proyecto/
├── go.mod
├── main.go
└── saludos/
    └── saludos.go
```

Archivo saludos.go:

```go
package saludos

func Hola(nombre string) string {
    return "¡Hola, " + nombre + "!"
}
```

Usar el paquete en main.go:

```go
package main

import (
    "fmt"
    "ejemplo.com/mimodulo/saludos"
)

func main() {
    mensaje := saludos.Hola("Mundo")
    fmt.Println(mensaje)
}
```

15. Manejo de errores
Error estándar:

go
Copiar código
import "errors"

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("no se puede dividir por cero")
    }
    return a / b, nil
}

resultado, err := dividir(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Resultado:", resultado)
}
16. Lectura y escritura de archivos
Escribir en un archivo:

go
Copiar código
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("archivo.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    file.WriteString("¡Hola, archivo!")
}
Leer de un archivo:

go
Copiar código
import (
    "fmt"
    "io/ioutil"
)

func main() {
    contenido, err := ioutil.ReadFile("archivo.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(contenido))
}
17. HTTP y servidores web
Servidor HTTP básico:

go
Copiar código
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, mundo!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}
Accede a http://localhost:8080 en tu navegador.

18. Buenas prácticas y herramientas
Formatear código: go fmt
Analizar código estático: go vet
Pruebas unitarias: Usa el paquete testing y ejecuta go test.
Documentación: Añade comentarios y genera documentación con godoc.
19. Recursos adicionales
Documentación oficial: https://golang.org/doc/
Tour de Go: https://tour.golang.org/welcome/1
Libros y cursos en línea: Hay numerosos recursos gratuitos y de pago para profundizar.
