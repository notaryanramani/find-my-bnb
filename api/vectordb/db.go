package vectordb

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type VectorDB struct {
	// Dimension of Vector
	dim int

	// Nodes in the VectorDB
	nodes []*Node
}

func NewVectorDB() *VectorDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dimStr, ok := os.LookupEnv("EMBEDDING_DIM")
	if !ok {
		log.Fatal("EMBEDDING_DIM is not set in .env file")
	}

	dim, err := strconv.Atoi(dimStr)
	if err != nil {
		log.Fatal(err)
	}

	return &VectorDB{
		dim:   dim,
		nodes: make([]*Node, 0),
	}
}

func (v *VectorDB) InitVectorDB(db *sql.DB) {
	// Steps
	// 1. Query all the nodes from the database
	tempNodes := queryDB(db)

	// 2. Get Embeddings for each node and add it to the VectorDB nodes
	for _, tempNode := range tempNodes {
		content := tempNode.Description + " " + tempNode.NeighborhoodOverview
		vector := getEmbeddings(content)
		v.AddNode(tempNode.ID, content, vector)
	}
}

func (v *VectorDB) AddNode(id int64, content string, vector []float64) {
	node := CreateNewNode(id, content, vector)
	v.nodes = append(v.nodes, node)
}

func (v *VectorDB) Size() int {
	return len(v.nodes)
}
