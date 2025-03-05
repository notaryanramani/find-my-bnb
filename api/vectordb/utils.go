package vectordb

import (
	"database/sql"
	"math/rand/v2"
)

func getRandomVector() []float64 {
	// Generate random vector for the content of 384 floats
	vector := make([]float64, 384)
	for i := 0; i < 384; i++ {
		vector[i] = rand.Float64()
	}
	return vector
}

func MultiplyVectors(v1 []float64, v2 []float64) float64 {
	v3 := 0.0
	for i := 0; i < len(v1); i++ {
		v3 += v1[i] * v2[i]
	}
	return v3
}

func AddElements(vector []float64) float64 {
	sum := 0.0
	for _, v := range vector {
		sum += v * v
	}
	return sum
}

func queryDB(db *sql.DB) *[]*DBNode {
	query := `SELECT id, description, neighborhood_overview FROM rooms`

	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	DBNodes := []*DBNode{}
	for rows.Next() {
		DBNode := &DBNode{}
		err := rows.Scan(
			&DBNode.ID,
			&DBNode.Description,
			&DBNode.NeighborhoodOverview,
		)
		if err != nil {
			return nil
		}

		DBNodes = append(DBNodes, DBNode)
	}

	return &DBNodes
}

func findNodeById(nodes []*Node, id int64) *Node {
	l := 0
	r := len(nodes) - 1

	for l <= r {
		m := l + (r-l)/2
		if nodes[m].ID == id {
			return nodes[m]
		}

		if nodes[m].ID < id {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return nil
}
