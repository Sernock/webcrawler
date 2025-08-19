package model

type Task struct {
	Id           string `json:"id"`
	SeedURL      string `json:"seed_url"`
	MaxPages     int    `json:"max_pages"`
	SameHostOnly bool   `json:"same_host_only"`
}

type Job struct {
	TaskID string `json:"struct_id"`
	URL    string `json:"url"`
}

type Result struct {
	TaskID string   `json:"task_id"`
	URL    string   `json:"url"`
	Titlte string   `json:"title,omitempty"`
	Status int      `json:"status"`
	Error  string   `json:"error,omitempty"`
	Links  []string `json:"links,omitempty"`
}

type Progress struct {
	TaskID       string `json:"task_id"`
	Enqueued     int    `json:"enqueued"`
	InFlight     int    `json:"in_flight"`
	Done         int    `json:"done"`
	Failed       int    `json:"failed"`
	WorkersAlive int    `json:"workers_alive"`
}
