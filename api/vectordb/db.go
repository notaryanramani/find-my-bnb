package vectordb

import (
	"database/sql"
	"encoding/gob"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type VectorDB struct {
	// Dimension of Vector
	Dim int

	// Nodes in the VectorDB
	Nodes []*Node

	// Embedder
	Embedder *Embedder
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
		Dim:      dim,
		Nodes:    make([]*Node, 0),
		Embedder: NewEmbedder(),
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

	// 3. Persist the VectorDB
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
	err = encoder.Encode(v)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadVectorDB() *VectorDB {
	var v VectorDB
	file, err := os.Open("persist/vectordb.gob")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatal(err)
		return nil
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
	gob.Register(&VectorDB{})
}
