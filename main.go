package main

import (
	"gfsd_go_ex1/common"
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

	studentRepository := student.NewStudentRepository()
	studentService := student.NewStudentService(dataSource, studentRepository)
	studentHandler := student.NewHandler(studentService)

	mux.HandleFunc("/students/{id}", studentHandler.GetStudent)
	mux.HandleFunc("/students", studentHandler.RegisterStudent)

	return mux
}
