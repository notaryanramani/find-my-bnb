package vectordb

import (
	"database/sql"
	"math"
)

type Node struct {
	// Node ID
	ID int64

	// Text content of the node
	Content string

	// Vector representation of the node
	Vector []float64
}

type DBNode struct {
	ID                   int64
	Description          sql.NullString
	NeighborhoodOverview sql.NullString
}

func CreateNewNode(id int64, content string, vector []float64) *Node {
	return &Node{
		ID:      id,
		Content: content,
		Vector:  vector,
	}
}

func (n *Node) Similarity(vector []float64) float64 {
	numerator := MultiplyVectors(n.Vector, vector)
	a_mod := math.Sqrt(AddElements(n.Vector))
	b_mod := math.Sqrt(AddElements(vector))
	return numerator / (a_mod * b_mod)
}
