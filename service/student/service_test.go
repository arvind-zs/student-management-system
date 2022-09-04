package student

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"student-management-system/models"
	"student-management-system/store"

	"github.com/golang/mock/gomock"
)

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		reqData   models.Student
		expRes    models.Student
		expGetRes []models.Student
		expGetErr error
		expErr    error
	}{
		{desc: "success:valid details posted successfully", reqData: models.Student{
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "09-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}, expRes: models.Student{
			ID:               1,
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "09-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}},
		{desc: "failure:post  method return query error", reqData: models.Student{
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "10-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}, expRes: models.Student{}, expErr: errors.New("query error")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().Get(ctx).Return(tc.expGetRes, tc.expGetErr)
		mockStore.EXPECT().Post(ctx, &tc.reqData).Return(tc.expRes, tc.expErr)

		res, err := mock.Post(ctx, &tc.reqData)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGet_Err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		reqData   models.Student
		expGetRes []models.Student
		expRes    models.Student
		expGetErr error
		expErr    error
	}{
		{desc: "failure:student details already exists", reqData: models.Student{
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "10-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}, expGetRes: []models.Student{{
			ID:               1,
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "10-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}}, expErr: errors.New("student already exists")},
		{desc: "failure:get method will return error", reqData: models.Student{
			FirstName:        "arvind",
			LastName:         "yadav",
			Gender:           "M",
			Dob:              "10-10-2000",
			MotherTongue:     "Hindi",
			Nationality:      "indian",
			FatherName:       "Kailash",
			MotherName:       "Indrawati",
			ContactNumber:    7348761063,
			FatherOccupation: "agriculture",
			MotherOccupation: "housewife",
		}, expGetErr: errors.New("query error"), expErr: errors.New("query error")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().Get(ctx).Return(tc.expGetRes, tc.expGetErr)

		res, err := mock.Post(ctx, &tc.reqData)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestPost_BodyCheckErr1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc    string
		reqData models.Student
		expRes  models.Student
		expErr  error
	}{
		{desc: "failure:invalid first name", reqData: models.Student{
			FirstName:     "a12",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid first name")},
		{desc: "failure:invalid first name", reqData: models.Student{
			FirstName:     "",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid first name")},
		{desc: "failure:invalid last name", reqData: models.Student{
			FirstName:     "arvind",
			LastName:      "ya12",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid last name")},
		{desc: "failure:invalid dob", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "04-31-2008",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob ..strconv error", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "ab-31-2008",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob ..strconv error", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "04-ab-2008",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob ..strconv error", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "04-30-abc",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob.. month is not correct", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "13-30-2000",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob.. date is not correct", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "12-32-2000",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
		{desc: "failure:invalid dob.. year is not correct", reqData: models.Student{
			FirstName:     "arvind",
			Dob:           "3-24-999",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid dob")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		res, err := mock.Post(ctx, &tc.reqData)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestPost_BodyCheckErr2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc    string
		reqData models.Student
		expRes  models.Student
		expErr  error
	}{
		{desc: "failure:invalid gender", reqData: models.Student{
			FirstName:     "arvind",
			Gender:        "K",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid gender")},
		{desc: "failure:invalid mother tongue", reqData: models.Student{
			FirstName:     "arvind",
			MotherTongue:  "a12",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid mother tongue")},
		{desc: "failure:invalid nationality", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "India123",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid nationality")},
		{desc: "failure:invalid nationality", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid nationality")},
		{desc: "failure:invalid father name", reqData: models.Student{
			FirstName:     "arvind",
			FatherName:    "123",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid father name")},
		{desc: "failure:invalid mother name", reqData: models.Student{
			FirstName:     "arvind",
			MotherName:    "123",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid mother name")},
		{desc: "failure:invalid contact number", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 734876106,
		}, expErr: errors.New("invalid contact number")},
		{desc: "failure:invalid father occupation", reqData: models.Student{
			FirstName:        "arvind",
			FatherOccupation: "a12",
			Nationality:      "Indian",
			ContactNumber:    7348761063,
		}, expErr: errors.New("invalid father occupation")},
		{desc: "failure:invalid mother occupation", reqData: models.Student{
			FirstName:        "arvind",
			MotherOccupation: "a12",
			Nationality:      "Indian",
			ContactNumber:    7348761063,
		}, expErr: errors.New("invalid mother occupation")},
		{desc: "failure:invalid family income", reqData: models.Student{
			FirstName:     "arvind",
			FamilyIncome:  -123,
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expErr: errors.New("invalid family income")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		res, err := mock.Post(ctx, &tc.reqData)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}
