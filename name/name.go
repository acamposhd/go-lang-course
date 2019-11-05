package name

import "fmt"

// GetName obtiene y retorna el nombre de una persona
func GetName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}
