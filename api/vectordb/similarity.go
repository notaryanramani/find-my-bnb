package vectordb

import (
	"sort"
	"sync"
)

type Similarity struct {
	nodeId     int64
	similarity float64
}

func (v *VectorDB) SimilaritySearch(text string) []*Node {
	vector := v.Embedder.getEmbeddings(text)
	similarities := make([]Similarity, len(v.Nodes))

	var wg sync.WaitGroup
	for i, node := range v.Nodes {
		wg.Add(1)
		go func(i int, node *Node) {
			defer wg.Done()
			similarities[i] = Similarity{
				nodeId:     node.ID,
				similarity: node.Similarity(vector),
			}
		}(i, node)
	}
	wg.Wait()

	// Sort the similarities in descending order
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].similarity > similarities[j].similarity
	})

	nodes := make([]*Node, len(v.Nodes))
	for i, similarity := range similarities {
		nodes[i] = v.Nodes[similarity.nodeId]
	}
	return nodes
}
