package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(Calculo(5,46))
	fmt.Println(Calculo(1,2,3,4))
	fmt.Println(Calculo(2))
	fmt.Println(Calculo(7,8,9,10,11))
}

/*retornar un valor entero
func uno(numero int) int {
	return numero * 2
}*/

//retornar un valor entero variante
func uno(numero int) (z int) {
	z=numero * 2
	return 
}

//retornar dos valores enteros
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

//parametros variables
func Calculo(numero ...int) int{
	total := 0
	for item , num:= range numero{
		total = total + num
		fmt.Printf("Item %d\n",item)
	}
	return total
}