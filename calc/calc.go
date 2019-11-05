package main

import (
	fmt "fmt"
)

func main() {
	menu()
}

func suma(num1 float32, num2 float32) float32 {
	return (num1 + num2)
}

func resta(num1 float32, num2 float32) float32 {
	return (num1 - num2)
}

func multiplicacion(num1 float32, num2 float32) float32 {
	return (num1 * num2)
}

func division(num1 float32, num2 float32) float32 {
	return (num1 / num2)
}

func opciones() int {
	var opcion int
	fmt.Println("Elige una opcion")
	fmt.Println("1.- Suma")
	fmt.Println("2.- Resta")
	fmt.Println("3.- Multiplicacion")
	fmt.Println("4.- Division")
	fmt.Println("0.- exit")
	fmt.Scanf("%d\n", &opcion)
	return opcion
}

func getNumeros() (float32, float32) {
	var num1, num2 float32
	fmt.Print("Ingresa el primer numero")
	fmt.Scanf("%f\n", &num1)
	fmt.Print("Ingresa el segundo numero")
	fmt.Scanf("%f\n", &num2)
	return num1, num2
}

func menu() {
	opcion := opciones()
	for opcion > 0 {

		fmt.Println()
		num1, num2 := getNumeros()
		if opcion == 1 {
			fmt.Println("La suma es: ", suma(num1, num2))
		} else if opcion == 2 {
			fmt.Println("La resta es: ", resta(num1, num2))
		} else if opcion == 3 {
			fmt.Println("La multiplicacion es: ", multiplicacion(num1, num2))
		} else if opcion == 4 {
			fmt.Println("La division es: ", division(num1, num2))
		}
		fmt.Println()
		opcion = opciones()

	}
}
