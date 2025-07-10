package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string = "Pedro"
var Estado bool = true
var Sueldo float32 = 1577.66
var Fecha time.Time = time.Now()

// Nombre = "Juan" // You can assign a new value to Nombre inside a function

func RestoVariables() {
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
