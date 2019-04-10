package main

import (
	"fmt"
	"math/rand"
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

	for i := range s {

		//fmt.Println(i)
		array := s[i]
		for j := range array {
			array[j] = rand.Float64()
		}

	}

	return s
}
