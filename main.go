package main

import (
	"gfsd_go_ex1/common"
	"gfsd_go_ex1/score"
	"gfsd_go_ex1/student"
	mux "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dataSource := common.NewDataSource()
	r := initializeMux(dataSource)
	log.Println("Listening on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initializeMux(dataSource *common.DataSource) *mux.Router {
	r := mux.NewRouter()

	// 직접 DI 를 해야만 한다.
	studentRepository := student.NewStudentRepository()
	studentService := student.NewStudentService(dataSource, studentRepository)
	studentHandler := student.NewHandler(studentService)

	scoreRepository := score.NewScoreRepository()
	scoreService := score.NewScoreService(dataSource, scoreRepository)
	scoreHandler := score.NewHandler(scoreService)

	// Syntactic sugar
	r.HandleFunc("/students/{studentId}", studentHandler.GetStudent).Methods(http.MethodGet)
	r.HandleFunc("/students", studentHandler.RegisterStudent).Methods(http.MethodPost)
	r.HandleFunc("/students/{studentId}/scores", scoreHandler.GetScores).Methods(http.MethodGet)
	r.HandleFunc("/students/{studentId}/scores", scoreHandler.RegisterScore).Methods(http.MethodPost)

	return r
}
