package vectordb

import (
	"sort"
)

type VectorDB struct {
	// Dimension of Vector
	dim int

	// Nodes in the VectorDB
	nodes []*Node
}

func InitVectorDB(dim int) *VectorDB {
	return &VectorDB{
		dim:   dim,
		nodes: make([]*Node, 0),
	}
}

func (v *VectorDB) AddNode(content string) {
	id := len(v.nodes)
	vector := getRandomVector() // Change this with a real vector, using HuggingFace embeddings model
	dbID := len(v.nodes)
	node := CreateNewNode(id, content, vector, dbID)

	v.nodes = append(v.nodes, node)
}

func (v *VectorDB) Size() int {
	return len(v.nodes)
}

type Similarity struct {
	nodeId     int
	similarity float64
}

func (v *VectorDB) SimilaritySearch(vector []float64) []*Node {
	similarities := make([]Similarity, len(v.nodes))
	for i, node := range v.nodes {
		similarities[i] = Similarity{
			nodeId:     node.ID,
			similarity: node.Similarity(vector),
		}
	}

	// Sort the similarities in descending order
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].similarity > similarities[j].similarity
	})

	nodes := make([]*Node, len(v.nodes))
	for i, similarity := range similarities {
		nodes[i] = v.nodes[similarity.nodeId]
	}
	return nodes
}
