package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	studentService *Service
}

func NewHandler(studentService *Service) *Handler {
	return &Handler{studentService}
}

func (h Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	parsedId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return
	}
	student := h.studentService.GetStudent(parsedId)

	var studentResponse StudentResponse
	if student != nil {
		studentResponse = StudentResponse{student.ID, student.Name}
	}
	studentJson, err := json.Marshal(studentResponse)
	if err != nil {
		http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, string(studentJson))
}

func (h Handler) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var studentResponse StudentResponse
		err := json.NewDecoder(r.Body).Decode(&studentResponse)
		if err != nil {
			http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
			return
		}
		studentJson, err := json.Marshal(studentResponse)
		if err != nil {
			http.Error(w, `{"error":"true"}`, http.StatusBadRequest)
			return
		}

		h.studentService.RegisterStudent(r.Context(), &Student{studentResponse.Id, studentResponse.Name})

		fmt.Fprint(w, string(studentJson))

	case http.MethodGet:
		http.Error(w, `{"error":"true"}`, http.StatusMethodNotAllowed)

	}

}
