package score

type Score struct {
	ID        int64
	Semester  int
	StudentId int64
	Score     int
}

type ScoreResponse struct {
	Id        int64 `json:"id"`
	Semester  int
	StudentId int64
	Score     int
}

type ScoreRequest struct {
	Id        int64
	Semester  int
	StudentId int64
	Score     int
}
