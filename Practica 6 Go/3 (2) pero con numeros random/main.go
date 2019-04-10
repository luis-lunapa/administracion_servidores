package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := CrearVector(10)

	fmt.Println(s)

}

func CrearVector(a int) []float64 {

	s := make([]float64, a)

	//random := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Float64()

	}

	return s
}
