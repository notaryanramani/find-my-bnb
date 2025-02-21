package vectordb

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Embedder struct {
	Url       string
	Url_batch string
}

type EmbeddingsRequest struct {
	Text string `json:"text"`
}

type EmbeddingsResponse struct {
	Embeddings []float64 `json:"embedding"`
}

func NewEmbedder() *Embedder {
	e, ok := os.LookupEnv("EMBEDDING_API_PORT")
	if !ok {
		log.Fatal("EMBEDDING_API_PORT is not set in .env file")
		return nil
	}

	return &Embedder{
		Url:       "http://127.0.0.1:" + e + "/embed",
		Url_batch: "http://127.0.0.1:" + e + "/embed_batch",
	}
}

func (e *Embedder) getEmbeddings(text string) []float64 {

	payload := EmbeddingsRequest{Text: text}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	resp, err := http.Post(e.Url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer resp.Body.Close()

	var embeddings EmbeddingsResponse
	err = json.NewDecoder(resp.Body).Decode(&embeddings)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return embeddings.Embeddings
}
