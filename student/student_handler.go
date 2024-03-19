package student

import (
	"encoding/json"
	"fmt"
	"gfsd_go_ex1/common/http/mime"
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
	// 처리중에 Panic 이 발생하면 이를 처리하는 코드를 Handler 에 설정합니다.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	w.Header().Add(mime.HeadContentType, mime.ContentTypeJson)

	idParam := r.PathValue("studentId")
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
	// 처리중에 Panic 이 발생하면 이를 처리하는 코드를 Handler 에 설정합니다.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	w.Header().Add(mime.HeadContentType, mime.ContentTypeJson)

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
