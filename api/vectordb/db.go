package vectordb

import (
	"database/sql"
	"encoding/gob"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
)

type VectorDB struct {
	// Dimension of Vector
	Dim int

	// Nodes in the VectorDB
	Nodes []*Node

	// Embedder
	Embedder *Embedder

	// Result Cache
	ResultCache map[string][]Similarity

	// Mutex
	Mu *sync.RWMutex
}

type VectorSearchRequest struct {
	Text    string `json:"text"`
	K       int    `json:"k"`
	Offset  int    `json:"offset"`
	QueryID string `json:"query_id"`
}

func NewVectorDB() *VectorDB {
	dimStr, ok := os.LookupEnv("EMBEDDING_DIM")
	if !ok {
		log.Fatal("EMBEDDING_DIM is not set in .env file")
	}

	dim, err := strconv.Atoi(dimStr)
	if err != nil {
		log.Fatal(err)
	}

	return &VectorDB{
		Dim:         dim,
		Nodes:       make([]*Node, 0),
		Embedder:    NewEmbedder(),
		ResultCache: make(map[string][]Similarity),
		Mu:          &sync.RWMutex{},
	}
}

func (v *VectorDB) InitVectorDB(db *sql.DB) {
	// Steps
	// 1. Query all the nodes from the database
	DBNodes := queryDB(db)

	// 2. Get Embeddings for each node and add it to the VectorDB nodes
	for _, DBNode := range *DBNodes {
		content := DBNode.Description.String + " " + DBNode.NeighborhoodOverview.String
		vector := v.Embedder.getEmbeddings(content)
		v.AddNode(DBNode.ID, content, vector)
	}

	// 3. Sort Nodes on IDs for Binary Search
	sort.Slice(v.Nodes, func(i, j int) bool {
		return v.Nodes[i].ID < v.Nodes[j].ID
	})

	// 4. Persist the VectorDB
	v.Persist()
}

func (v *VectorDB) Persist() {
	err := os.MkdirAll("persist", 0755)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("persist/vectordb.gob")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(v.Nodes)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadVectorDB() *VectorDB {
	var n []*Node
	file, err := os.Open("persist/vectordb.gob")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&n)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	v := VectorDB{
		Dim: len((n)[0].Vector),
		Nodes: n,
		Embedder: NewEmbedder(),
		ResultCache: make(map[string][]Similarity),
		Mu: &sync.RWMutex{},
	}

	return &v
}

func (v *VectorDB) AddNode(id int64, content string, vector []float64) {
	node := CreateNewNode(id, content, vector)
	v.Nodes = append(v.Nodes, node)
}

func (v *VectorDB) Size() int {
	return len(v.Nodes)
}

func init() {
	gob.Register(&[]*Node{})
}
