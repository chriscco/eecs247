package entity 

type Response struct {
	Cnts []WordCount
}

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}