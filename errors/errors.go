package main

import (
	"errors"
	"fmt"
)

func main() {
	number, err := Sum(50, 50)
	if err != nil {
		panic(err)
	}
	fmt.Println(number)
}

func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}
