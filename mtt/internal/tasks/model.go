package tasks

type TaskRequest struct {
	Task string `json:"task"`
}

type TaskBody struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}
