package model

type Todo struct {
	ID          int64   `json: "id"`
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Finish      bool    `json: "finish"`
	Position    float32 `json: "position"`
}
