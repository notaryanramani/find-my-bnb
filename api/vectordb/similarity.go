package vectordb

import (
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Similarity struct {
	nodeId     int64
	similarity float64
}

func (v *VectorDB) SimilaritySearch(req VectorSearchRequest) ([]*Node, string) {
	vector := v.Embedder.getEmbeddings(req.Text)
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

	queryId := uuid.New().String()

	v.Mu.Lock()
	v.ResultCache[queryId] = similarities
	v.Mu.Unlock()

	time.AfterFunc(time.Minute*10, func() {
		v.Mu.Lock()
		delete(v.ResultCache, queryId)
		v.Mu.Unlock()
	})

	nodes := make([]*Node, req.K)
	for i := range req.K {
		nodes[i] = v.Nodes[similarities[i].nodeId]
	}

	return nodes, queryId
}

func (v *VectorDB) GetNodesFromCache(req VectorSearchRequest) []*Node {
	v.Mu.RLock()
	defer v.Mu.RUnlock()

	cache, _ := v.ResultCache[req.QueryID]
	
	nodes := make([]*Node, req.K)
	for i := range req.K {
		nodes[i] = v.Nodes[cache[i+req.Offset].nodeId]
	}
	return nodes
}
