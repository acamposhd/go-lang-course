package maps

// GetMap devuelve un map
func GetMap() map[string]int {
	// mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Juan":   18,
		"Yohan":  24,
		"Freddy": 31,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100

	delete(mapTest, "llave1")
	delete(mapTest, "Yohan")
	return mapTest
}

// GetEdad devuelve la edad
func GetEdad(name string) int {
	dic1 := GetMap()
	return dic1[name]
}
