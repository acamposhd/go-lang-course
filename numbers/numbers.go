package numbers

// GetVariables devuelve 3 números enteros
func GetVariables() (int, int32, int64) {
	return 1, 2147000000, 903131313131311313
}

// GetFloat devuelve dos números de punto flotante
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum suma dos números enteros y devuelve el resultado
func Sum(a int, b int) int {
	return a + b
}
