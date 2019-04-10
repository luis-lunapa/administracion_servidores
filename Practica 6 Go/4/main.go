package main

import (
	"fmt"
)

func main() {

	s := CrearMatriz(10)

	fmt.Println(s)

}

func CrearMatriz(a int) [][]float64 {

	s := make([][]float64, a, a)
	for i := range s {
		s[i] = make([]float64, a)
	}

	return s
}
