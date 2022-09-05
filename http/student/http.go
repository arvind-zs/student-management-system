package student

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"student-management-system/models"
	"student-management-system/service"

	"github.com/gorilla/mux"
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

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	res, err := h.student.Get(r.Context(), firstName, lastName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println(err.Error())

		return
	}
}

func (h handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	student, err := h.student.GetByID(r.Context(), ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	body, err := json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println(err.Error())

		return
	}
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	err = h.student.Delete(r.Context(), ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())

			return
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
