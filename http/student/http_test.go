package student

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"student-management-system/models"
	"student-management-system/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockStudent(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc      string
		reqBody   models.Student
		expRes    models.Student
		expErr    error
		expStatus int
	}{
		{desc: "success:valid details posted successfully", reqBody: models.Student{
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expRes: models.Student{
			ID:            1,
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expStatus: http.StatusCreated},
		{desc: "failure:invalid details  ", reqBody: models.Student{
			FirstName:     "",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid first name"), expStatus: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			log.Println(err.Error())
		}

		req := httptest.NewRequest(http.MethodPost, "/student", bytes.NewReader(body))
		w := httptest.NewRecorder()

		mockService.EXPECT().Post(req.Context(), &tc.reqBody).Return(tc.expRes, tc.expErr)
		mock.Post(w, req)

		if w.Code != tc.expStatus {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expStatus, w.Code)
		}
	}
}

func TestPost_UnmarshallingError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockStudent(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc      string
		reqBody   []byte
		expRes    models.Student
		expErr    error
		expStatus int
	}{
		{desc: "failure:unmarshalling error", reqBody: []byte(`{
			FirstName:     arvind,
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}`), expErr: errors.New("invalid body"), expStatus: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "/student", bytes.NewReader(tc.reqBody))
		w := httptest.NewRecorder()

		mock.Post(w, req)

		if w.Code != tc.expStatus {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expStatus, w.Code)
		}
	}
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockStudent(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc      string
		id        string
		expRes    models.Student
		expErr    error
		expStatus int
	}{
		{desc: "success:get details with valid id", id: "1", expRes: models.Student{
			ID:            1,
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expStatus: http.StatusOK},
		{desc: "failure: invalid id", id: "1111", expErr: errors.New("invalid id not present in db rows set"),
			expStatus: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/student/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id})

		ID, err := strconv.Atoi(tc.id)
		if err != nil {
			log.Println(err.Error())
		}

		mockService.EXPECT().GetByID(req.Context(), ID).Return(tc.expRes, tc.expErr)

		mock.GetByID(w, req)

		if w.Code != tc.expStatus {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expStatus, w.Code)
		}
	}
}

func TestGetByID_StrConvErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockStudent(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc      string
		id        string
		expRes    models.Student
		expErr    error
		expStatus int
	}{
		{desc: "failure: strconv error", id: "abc", expErr: errors.New("invalid id will give strconv error"),
			expStatus: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/student/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id})

		mock.GetByID(w, req)

		if w.Code != tc.expStatus {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expStatus, w.Code)
		}
	}
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockStudent(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expRes    []models.Student
		expErr    error
		expStatus int
	}{
		{desc: "success:valid query params firstName and lastName", firstName: "arvind", lastName: "yadav", expRes: []models.Student{
			{ID: 1, FirstName: "arvind", LastName: "yadav", Nationality: "Indian", ContactNumber: 7348761063},
		}, expStatus: http.StatusOK},
		{desc: "success:valid query params firstName ", firstName: "anuj", expRes: []models.Student{
			{ID: 1, FirstName: "arvind", Nationality: "Indian", ContactNumber: 1234567891},
		}, expStatus: http.StatusOK},
		{desc: "success:valid query params lastName ", lastName: "bhagat", expRes: []models.Student{
			{ID: 1, FirstName: "chetan", LastName: "bhagat", Nationality: "Indian", ContactNumber: 1234567875},
		}, expStatus: http.StatusOK},
		{desc: "failure:invalid query params", firstName: "12345", lastName: "5678",
			expErr: errors.New("no rows in db result set"), expStatus: http.StatusBadRequest},
		{desc: "failure:invalid query params", firstName: "", lastName: "",
			expErr: errors.New("missing  both query params"), expStatus: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/student?"+"firstName="+tc.firstName+"&"+"lastName="+tc.lastName, nil)

		mockService.EXPECT().Get(req.Context(), tc.firstName, tc.lastName).Return(tc.expRes, tc.expErr)

		mock.Get(w, req)

		if w.Code != tc.expStatus {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expStatus, w.Code)
		}
	}
}
