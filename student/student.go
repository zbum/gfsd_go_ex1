package student

type StudentResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Student struct {
	ID   int64
	Name string
}
