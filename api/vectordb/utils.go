package vectordb

import (
	"math/rand/v2"
)

func getRandomVector() []float64 {
	// Generate random vector for the content of 300 floats
	vector := make([]float64, 300)
	for i := 0; i < 300; i++ {
		vector[i] = rand.Float64()
	}
	return vector
}


func MultiplyVectors(v1 []float64, v2 []float64) float64 {
	v3 := 0.0
	for i := 0; i < len(v1); i++ {
		v3 += v1[i] * v2[i]
	}
	return v3
}

func AddElements(vector []float64) float64 {
	sum := 0.0
	for _, v := range vector {
		sum += v
	}
	return sum
}