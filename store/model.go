package store

type Todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type Data struct {
	Todos []Todo `json:"todos"`
}
