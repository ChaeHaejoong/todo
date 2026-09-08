package store

type Todo struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Date    string `json:"date,omitempty"`
	Time    string `json:"time,omitempty"`
}

type Appointment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Date      string `json:"date,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

type Data struct {
	Todos       []Todo        `json:"todos"`
	Appointment []Appointment `json:"Appointment"`
}
