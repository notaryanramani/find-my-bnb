package vectordb

import (
	"sort"
	"sync"
)

type Similarity struct {
	nodeId     int64
	similarity float64
}

func (v *VectorDB) SimilaritySearch(text string, topK int) []*Node {
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

	nodes := make([]*Node, topK)
	for i := 0; i < topK; i++ {
		nodes[i] = v.Nodes[similarities[i].nodeId]
	}
	return nodes
}
