# Lenguaje Golang
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Repositorio para aprender conceptos del lenguaje Go siguiendo algunos tutorial de Youtube:
- [Tiago](https://www.youtube.com/watch?v=h3fqD6IprIA&list=WL&index=1&t=1098s)
- [Nana](https://www.youtube.com/watch?v=yyUHQIec83I)
- [Nana Crash Crouse](https://www.youtube.com/watch?v=XCZWyN9ZbEQ&t=42s)
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

### 15. Manejo de errores
Error estándar:

```go
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
```


### 16. Lectura y escritura de archivos
Escribir en un archivo:

```go

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
```

Leer de un archivo:

```go
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
```

### 17. HTTP y servidores web
Servidor HTTP básico:

```go

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
```
Accede a http://localhost:8080 en tu navegador.

### 18. Buenas prácticas y herramientas
- Formatear código: go fmt
- Analizar código estático: go vet
- Pruebas unitarias: Usa el paquete testing y ejecuta go test.
- Documentación: Añade comentarios y genera documentación con godoc.


### 19. Recursos adicionales
- Documentación oficial: https://golang.org/doc/
- Tour de Go: https://tour.golang.org/welcome/1
- Libros y cursos en línea: Hay numerosos recursos gratuitos y de pago para profundizar.


## Keywords
`break`

Termina de inmediato el bucle o la sentencia switch en la que se encuentra.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
// Imprime números del 0 al 4 y luego se detiene.
```

`case`

Especifica una condición en una sentencia switch o select para comparar valores.

```go
switch day := "lunes"; day {
case "lunes":
    fmt.Println("Inicio de semana")
case "viernes":
    fmt.Println("Fin de semana")
default:
    fmt.Println("Día intermedio")
}
// Salida: Inicio de semana
```

`chan`

Declara un canal, que es una vía para la comunicación entre goroutines.

```go
c := make(chan string)

go func() {
    c <- "Mensaje desde la goroutine"
}()
msg := <-c
fmt.Println(msg)
// Salida: Mensaje desde la goroutine
```

`const`

Declara una constante, un valor que no puede ser modificado después de su definición.

```go
const Pi = 3.1416
fmt.Println("El valor de Pi es:", Pi)
// Salida: El valor de Pi es: 3.1416
```


`continue`

Salta a la siguiente iteración del bucle, omitiendo el resto del código en la iteración actual.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
// Imprime 0, 1, 3, 4 (omite el 2)
```


`default`

Especifica el bloque de código a ejecutar en una sentencia switch o select si no hay coincidencia con los casos anteriores.

```go
color := "verde"
switch color {
case "rojo":
    fmt.Println("El color es rojo")
case "azul":
    fmt.Println("El color es azul")
default:
    fmt.Println("Color no reconocido")
}
// Salida: Color no reconocido
```

`defer`

Retrasa la ejecución de una función hasta que la función que la contiene finaliza.

```go
func main() {
    defer fmt.Println("Esto se ejecuta al final")
    fmt.Println("Esto se ejecuta primero")
}
// Salida:
// Esto se ejecuta primero
// Esto se ejecuta al final
```

`else`

Especifica el bloque de código a ejecutar si la condición en un if es falsa.

```go
edad := 16
if edad >= 18 {
    fmt.Println("Eres mayor de edad")
} else {
    fmt.Println("Eres menor de edad")
}
// Salida: Eres menor de edad
```

`fallthrough`

En una sentencia switch, permite que la ejecución continúe al siguiente caso sin evaluar su condición.

```go
number := 1
switch number {
case 1:
    fmt.Println("Uno")
    fallthrough
case 2:
    fmt.Println("Dos")
default:
    fmt.Println("Otro número")
}
// Salida:
// Uno
// Dos
```

`for`

Inicia un bucle, que puede ser con una condición, con una cláusula de inicialización y post, o un bucle infinito.

```go
// Bucle clásico
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

// Bucle como 'while'
i := 0
for i < 3 {
    fmt.Println(i)
    i++
}

// Bucle infinito
for {
    fmt.Println("Bucle infinito")
    break // Necesario para salir del bucle
}
```

`func`

Declara una función o un método.

```go
func suma(a int, b int) int {
    return a + b
}

resultado := suma(3, 4)
fmt.Println("Resultado:", resultado)
// Salida: Resultado: 7
```

`go` 

Inicia una goroutine, es decir, ejecuta una función de forma concurrente.

```go
func decirHola() {
    fmt.Println("Hola")
}

go decirHola()
fmt.Println("Mundo")
// Salida:
// Mundo
// Hola
// (El orden puede variar debido a la concurrencia)
```

`goto`

Transfiere el control de ejecución a una etiqueta específica dentro de la función actual.

```go
func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            goto Salir
        }
        fmt.Println(i)
    }
Salir:
    fmt.Println("Se utilizó goto para salir del bucle")
}
// Salida:
// 0
// 1
// 2
// Se utilizó goto para salir del bucle
```

`if`

Inicia una estructura condicional que ejecuta un bloque de código si la condición es verdadera.

```go
temperatura := 30
if temperatura > 25 {
    fmt.Println("Hace calor")
}
// Salida: Hace calor
```

`import`

Importa paquetes para utilizar sus funciones y tipos en el programa actual.

```go
import (
    "fmt"
    "math"
)

func main() {
    raiz := math.Sqrt(16)
    fmt.Println("La raíz cuadrada de 16 es:", raiz)
}
`
// Salida: La raíz cuadrada de 16 es: 4

`interface`

Declara una interfaz, un conjunto de métodos que un tipo puede implementar.

```go
type Forma interface {
    Area() float64
}

type Cuadrado struct {
    Lado float64
}

func (c Cuadrado) Area() float64 {
    return c.Lado * c.Lado
}

func main() {
    var f Forma = Cuadrado{Lado: 5}
    fmt.Println("Área del cuadrado:", f.Area())
}
// Salida: Área del cuadrado: 25
```

`map`

Declara un mapa, una estructura de datos que almacena pares clave-valor desordenados.

```go
capitales := map[string]string{
    "España":     "Madrid",
    "Francia":    "París",
    "Argentina":  "Buenos Aires",
}

fmt.Println("La capital de Francia es:", capitales["Francia"])
// Salida: La capital de Francia es: París
```


`package`

Especifica el paquete al que pertenece el archivo fuente.

```go
// En el archivo main.go
package main

import "fmt"

func main() {
    fmt.Println("Programa principal")
}
```

`range`

Itera sobre elementos en una variedad de estructuras de datos como arrays, slices, mapas y canales.

```go
numeros := []int{10, 20, 30}
for indice, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
}
// Salida:
// Índice: 0, Valor: 10
// Índice: 1, Valor: 20
// Índice: 2, Valor: 30
```


`return`

Finaliza la ejecución de una función y opcionalmente devuelve uno o más valores.

```go
func doble(n int) int {
    return n * 2
}

resultado := doble(4)
fmt.Println("El doble es:", resultado)
// Salida: El doble es: 8
```

`select`

Espera en múltiples operaciones de envío o recepción en canales y ejecuta el caso que esté listo.

```go
c1 := make(chan string)
c2 := make(chan string)

go func() {
    c1 <- "Mensaje 1"
}()

go func() {
    c2 <- "Mensaje 2"
}()

select {
case msg1 := <-c1:
    fmt.Println(msg1)
case msg2 := <-c2:
    fmt.Println(msg2)
}
// Salida: Puede ser "Mensaje 1" o "Mensaje 2" dependiendo de cuál canal esté listo primero.
```

`struct`

Declara una estructura, un tipo de datos compuesto que agrupa campos bajo un mismo nombre.

```go
type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    p := Persona{Nombre: "Lucas", Edad: 30}
    fmt.Println("Nombre:", p.Nombre)
    fmt.Println("Edad:", p.Edad)
}
// Salida:
// Nombre: Lucas
// Edad: 30
```

`switch`

Inicia una estructura de control que selecciona entre múltiples casos basados en el valor de una expresión.

```go
nota := 85
switch {
case nota >= 90:
    fmt.Println("Excelente")
case nota >= 80:
    fmt.Println("Muy bien")
case nota >= 70:
    fmt.Println("Bien")
default:
    fmt.Println("Necesita mejorar")
}
// Salida: Muy bien
```

`type`

Declara un nuevo tipo de dato o define un alias para un tipo existente.

```go
type Texto string

func main() {
    var mensaje Texto = "Hola, mundo"
    fmt.Println(mensaje)
}
// Salida: Hola, mundo
```

`var`

Declara una o más variables, opcionalmente con valores iniciales.

```go
var nombre string = "Ana"
var edad int
edad = 28

fmt.Println("Nombre:", nombre)
fmt.Println("Edad:", edad)
// Salida:
// Nombre: Ana
// Edad: 28
```