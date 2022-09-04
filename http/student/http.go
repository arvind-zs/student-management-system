package student

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"student-management-system/models"
	"student-management-system/service"
)

type handler struct {
	student service.Student
}

func New(s service.Student) handler {
	return handler{student: s}
}

func (h handler) Post(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	var student models.Student

	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	student, err = h.student.Post(r.Context(), &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	body, err = json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(body)
	if err != nil {
		log.Println(err.Error())

		return
	}
}
