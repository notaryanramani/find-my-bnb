package vectordb

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestDBInit(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Errorf("Error loading .env file")
	}

	db := NewVectorDB()
	if db == nil {
		t.Errorf("DB not initialized")
	}
}

func TestDBAddNode(t *testing.T) {
	db := NewVectorDB()
	vector := getRandomVector()
	db.AddNode(1, "test", vector)
	if db.Size() != 1 {
		t.Errorf("Node insertion failed, expected 1, got %d\n", db.Size())
	}
}

func TestNodeSimilarity(t *testing.T) {
	vector1 := getRandomVector()
	vector2 := getRandomVector()
	node1 := CreateNewNode(1, "test", vector1)
	similarity := node1.Similarity(vector2)

	if similarity == 0 {
		t.Errorf("Similarity calculation failed")
	}

	fmt.Printf("Similarity: %f\n", similarity)
}

func TestMultiplyVectors(t *testing.T) {
	v1 := []float64{1, 2, 3}
	v2 := []float64{4, 5, 6}
	result := MultiplyVectors(v1, v2)
	if result != 32 {
		t.Errorf("Multiplication failed, expected 32, got %f\n", result)
	}
}
