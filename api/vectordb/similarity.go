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
	vector := getEmbeddings(text)
	similarities := make([]Similarity, len(v.nodes))

	var wg sync.WaitGroup
	for i, node := range v.nodes {
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

	nodes := make([]*Node, len(v.nodes))
	for i, similarity := range similarities {
		nodes[i] = v.nodes[similarity.nodeId]
	}
	return nodes
}
