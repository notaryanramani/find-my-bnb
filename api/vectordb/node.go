package vectordb

import (
	"math"
)

type Node struct {
	// Node ID
	ID int

	// Text content of the node
	content string

	// Vector representation of the node
	vector []float64

	// DB ID
	dbID int
}

func CreateNewNode(id int, content string, vector []float64, dbID int) *Node {
	return &Node{
		ID:      id,
		content: content,
		vector:  vector,
		dbID:    dbID,
	}
}

func (n *Node) Similarity (vector []float64) float64 {
	numerator := MultiplyVectors(n.vector, vector)
	a_mod := math.Sqrt(AddElements(n.vector))
	b_mod := math.Sqrt(AddElements(vector))
	return numerator / (a_mod * b_mod)
}
