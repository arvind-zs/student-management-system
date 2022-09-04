package student

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"student-management-system/models"
	"student-management-system/service"

	"github.com/golang/mock/gomock"
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

		mockService.EXPECT().Post(req.Context(), tc.reqBody).Return(tc.expRes, tc.expErr)
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
