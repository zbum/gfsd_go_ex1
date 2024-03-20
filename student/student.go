package student

type StudentRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type StudentResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Student struct {
	ID   int64
	Name string
}
