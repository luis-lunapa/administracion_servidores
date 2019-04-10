package main

import (
	"fmt"
	"math/rand"
)

func main() {

	array := CrearVector(10)
	matrix := CrearMatriz(10)

	result := MultVectorMatriz(array, matrix)

	fmt.Println(result)

}

func MultVectorMatriz(array []float64, matrix [][]float64) [][]float64 {

	resultado := make([][]float64, len(matrix))
	for i := range resultado {
		resultado[i] = make([]float64, len(matrix))
	}

	for i := range matrix {
		columna := matrix[i]
		for j := range columna {
			//mt.Println(columna[j] * array[j])
			resultado[i][j] = columna[j] * array[j]

		}

	}

	return resultado
}

func CrearVector(a int) []float64 {

	s := make([]float64, a)

	//random := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Float64()

	}

	return s
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
