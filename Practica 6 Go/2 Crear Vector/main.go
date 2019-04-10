package main

import "fmt"

func main() {

	s := CrearVector(10)

	fmt.Println(s)

}

func CrearVector(a int) []float64 {
	return make([]float64, a)
}
