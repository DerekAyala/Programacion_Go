package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	/* BREAK
	numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}
	fmt.Println("termine")
	*/

	/* CONTINUE
	var i = 0
	for i <= 10{
		fmt.Printf("\n valor de i: %d",i)
		if i == 5{
			fmt.Printf("\n multiplicamors por 2 ")
			i=i*2
			continue
		}
		fmt.Printf("\tPaso por aqui\n")
		i++
	}
	*/

	/*	GOTO
		var i int = 0

		RUTINA:
			for i < 10{
				if i==4{
					i=i+2
					fmt.Println("voy a RUTINA 2 A I")
					goto RUTINA
				}
				fmt.Printf("Valor de i: %d\n",i)
				i++
			}
	*/
}
