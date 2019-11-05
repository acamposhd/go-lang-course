package main

import (
	"fmt"
	"strings"

	"./flow"
	"./maps"
	"./name"
	"./numbers"
	"./structs"
)

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	lastname := "<Modificar con mi apellido>"
	number := numbers.Sum(50, 50)
	firstName := name.GetName()
	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()
	stringUTF8 := name.GetUnicode()
	fmt.Printf(helloWorld, firstName, lastname)
	fmt.Print("Hola mundo")
	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Cadena con UTF8:", stringUTF8)
	fmt.Println(string("hola"[1]))
	fmt.Println("Cantidad de letras que tiene hola ", len("hola"))
	structs.GetArray()
	structs.GetSlice()
	flow.IfTest()
	strings2()
	flow.SwitchTest()
	maps.GetEdad("Juan")
}

func strings2() {
	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))
}
