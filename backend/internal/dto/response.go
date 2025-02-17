package dto

type Response struct {
	Message string  `json:"message,omitempty"`
	Data    any     `json:"data,omitempty"`
	Errors  []Error `json:"errors,omitempty"`
}

type Error struct {
	Field  string `json:"field"`
	Detail string `json:"detail"`
}