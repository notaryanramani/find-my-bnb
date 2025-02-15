package vectordb

type Similarity struct {
	nodeId int
	similarity float64
}

type BySimilarity []Similarity

func (s BySimilarity) Len() int           { return len(s) }
func (s BySimilarity) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySimilarity) Less(i, j int) bool { return s[i].similarity < s[j].similarity }