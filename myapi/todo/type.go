package todo

type Todo struct {
	ID     string `json:"identifier"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
