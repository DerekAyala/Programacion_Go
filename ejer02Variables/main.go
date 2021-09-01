package main

import (
	"fmt"
	"strconv"
)

/*crear variable
var numero int
var texto string
*/var status bool

//float32 float64

func main() {
	//crear varias variables
	var numero, numero2 int
	numero, numero2, texto, status := 3, 2, "hola", true
	//asignar variables
	numero3 := 4
	numero3 = 15
	var moneda float32 = 0
	//casting
	numero2 = int(moneda)
	//texto = fmt.Sprintf("%f", moneda)
	texto = strconv.Itoa(int(moneda))
	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
	fmt.Println(status)
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
