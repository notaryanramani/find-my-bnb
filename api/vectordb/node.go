package vectordb

import (
	"math"
)

type Node struct {
	// Node ID
	ID int64

	// Text content of the node
	content string

	// Vector representation of the node
	vector []float64
}

type DBNode struct {
	ID                   int64
	Description          string
	NeighborhoodOverview string
}

func CreateNewNode(id int64, content string, vector []float64) *Node {
	return &Node{
		ID:      id,
		content: content,
		vector:  vector,
	}
}

func (n *Node) Similarity(vector []float64) float64 {
	numerator := MultiplyVectors(n.vector, vector)
	a_mod := math.Sqrt(AddElements(n.vector))
	b_mod := math.Sqrt(AddElements(vector))
	return numerator / (a_mod * b_mod)
}
