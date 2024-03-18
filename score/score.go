package score

type Score struct {
	ID        int64
	Semester  int
	StudentId int64
	Score     int
}

type ScoreResponse struct {
	Id        int64 `json:"id"`
	Semester  int   `json:"semester"`
	StudentId int64 `json:"studentId"`
	Score     int   `json:"score"`
}

type ScoreRequest struct {
	Id        int64 `json:"id"`
	Semester  int   `json:"semester"`
	StudentId int64 `json:"studentId"`
	Score     int   `json:"score"`
}
