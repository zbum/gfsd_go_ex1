package main

import (
	"gfsd_go_ex1/common"
	"gfsd_go_ex1/score"
	"gfsd_go_ex1/student"
	"log"
	"net/http"
)

func main() {
	dataSource := common.NewDataSource()
	mux := initializeMux(dataSource)
	log.Println("Listening on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func initializeMux(dataSource *common.DataSource) *http.ServeMux {
	mux := http.NewServeMux()

	// 직접 DI 를 해야만 한다.
	studentRepository := student.NewStudentRepository()
	studentService := student.NewStudentService(dataSource, studentRepository)
	studentHandler := student.NewHandler(studentService)

	scoreRepository := score.NewScoreRepository()
	scoreService := score.NewScoreService(dataSource, scoreRepository)
	scoreHandler := score.NewHandler(scoreService)

	mux.HandleFunc("/students/{studentId}", studentHandler.GetStudent)
	mux.HandleFunc("/students", studentHandler.RegisterStudent)

	mux.HandleFunc("/students/{studentId}/scores", scoreHandler.ProcessScores)

	return mux
}
