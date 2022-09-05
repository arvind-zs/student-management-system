package main

import (
	"fmt"
	"log"
	"net/http"

	"student-management-system/driver"
	student3 "student-management-system/http/student"
	student2 "student-management-system/service/student"
	"student-management-system/store/student"

	"github.com/gorilla/mux"
)

func main() {
	db, err := driver.Connection()
	if err != nil {
		log.Println(err.Error())
	}

	defer db.Close()

	//   injecting dependencies
	storeStudent := student.New(db)
	serviceStudent := student2.New(storeStudent)
	handlerStudent := student3.New(serviceStudent)

	r := mux.NewRouter()
	r.HandleFunc("/student", handlerStudent.Post).Methods(http.MethodPost)
	r.HandleFunc("/student/{id}", handlerStudent.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/student", handlerStudent.Get).Methods(http.MethodGet)

	fmt.Println("http server started and listening on port :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
