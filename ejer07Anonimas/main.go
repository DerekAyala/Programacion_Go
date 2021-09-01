package main

import "fmt"

var Calculo func(int,int) int = func(num1 int,num2 int) int {
	return num1 + num2
}
func main(){
	fmt.Printf("Sumo 5 + 7 = %d\n",Calculo(5,7))
	//restamos
	Calculo = func(num1 int , num2 int) int{
		return num1-num2
	}
	fmt.Printf("resta 6 - 4 = %d\n",Calculo(6,4))
	//dividir
	Calculo = func(num1 int , num2 int) int{
		return num1/num2
	}
	fmt.Printf("divido 18 / 6 = %d\n",Calculo(18,6))
	Operaciones()

	//CLOSURE
	tabla02:=2
	miTabla:=Tabla(tabla02)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}

func Operaciones(){
	resultado := func () int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}
//Closures
func Tabla(valor int) func() int {
	numero:=valor
	secuencia:=0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}