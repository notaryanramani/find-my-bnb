package vectordb

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"bytes"

	"github.com/joho/godotenv"
)

type EmbeddingsRequest struct {
	Text string `json:"text"`
}

type EmbeddingsResponse struct {
	Embeddings []float64 `json:"embedding"`
}

func getEmbeddings(text string) []float64 {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	e, ok := os.LookupEnv("EMBEDDING_PORT")
	if !ok {
		log.Fatal("EMBEDDING_API_PORT is not set in .env file")
		return nil
	}

	url := "http://localhost:" + e + "/embed"

	payload := EmbeddingsRequest{Text: text}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
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