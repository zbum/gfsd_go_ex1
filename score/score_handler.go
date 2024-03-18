package score

import (
	"encoding/json"
	"fmt"
	"gfsd_go_ex1/common/http/mime"
	"net/http"
	"strconv"
)

type Handler struct {
	scoreService *Service
}

func NewHandler(studentService *Service) *Handler {
	return &Handler{studentService}
}

func (h Handler) ProcessScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(mime.HeadContentType, mime.ContentTypeJson)

	switch r.Method {
	case http.MethodGet:
		h.getScores(w, r)
	case http.MethodPost:
		h.registerScore(w, r)
	}
}
func (h Handler) getScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(mime.HeadContentType, mime.ContentTypeJson)

	studentIdParam := r.PathValue("studentId")
	parsedStudentId, err := strconv.ParseInt(studentIdParam, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
		return
	}
	scores := h.scoreService.GetScores(r.Context(), parsedStudentId)

	var scoreResponses = make([]ScoreResponse, len(scores))
	if scores != nil {
		for _, score := range scores {
			scoreResponses = append(scoreResponses, ScoreResponse{score.ID, score.Semester, score.StudentId, score.Score})
		}
	}
	scoresJson, err := json.Marshal(scoreResponses)
	if err != nil {
		http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, string(scoresJson))
}

func (h Handler) registerScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(mime.HeadContentType, mime.ContentTypeJson)

	var scoreRequest ScoreRequest
	err := json.NewDecoder(r.Body).Decode(&scoreRequest)
	if err != nil {
		http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
		return
	}
	studentJson, err := json.Marshal(scoreRequest)
	if err != nil {
		http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
		return
	}

	h.scoreService.RegisterStudent(r.Context(), &Score{scoreRequest.Id, scoreRequest.Semester, scoreRequest.StudentId, scoreRequest.Score})

	fmt.Fprint(w, string(studentJson))

}
