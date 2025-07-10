package main

//import "fmt"
import (
	"fmt"

	"github.com/marcelo-a-ipar/GoDesdeCero/variables"
)

func main() {
	//	fmt.Println("¡Hola, mundo!")

	//variables.MuestroEnteros()

	//variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(12345)

	fmt.Println(estado)
	fmt.Println(texto)
}
