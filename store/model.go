package store

type Todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
}

type Data struct {
	Todos []Todo `json:"todos"`
}
