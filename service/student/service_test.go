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

func TestPost_GetErr(t *testing.T) {
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

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc   string
		id     int
		expRes models.Student
		expErr error
	}{
		{desc: "success:fetch student details with valid id", id: 1, expRes: models.Student{
			ID:            1,
			FirstName:     "arvind",
			LastName:      "yadav",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}},

		{desc: "failure: invalid id not present in result set", id: 1111, expErr: errors.New("no rows in db result set")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().GetByID(ctx, tc.id).Return(tc.expRes, tc.expErr)

		res, err := mock.GetByID(ctx, tc.id)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGet_FirstAndLastName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expRes    []models.Student
		expErr    error
	}{
		{desc: "success:valid Query params firstName and lastName", firstName: "arvind", lastName: "yadav",
			expRes: []models.Student{
				{ID: 1, FirstName: "arvind", LastName: "yadav", Nationality: "Indian", ContactNumber: 7348761063},
			}},
		{desc: "failure:invalid Query params firstName and lastName", firstName: "123a", lastName: "345a",
			expErr: errors.New("no rows present in database with this query params")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().GetByFirstAndLastName(ctx, tc.firstName, tc.lastName).Return(tc.expRes, tc.expErr)

		res, err := mock.Get(ctx, tc.firstName, tc.lastName)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGet_FirstName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expRes    []models.Student
		expErr    error
	}{
		{desc: "success:valid Query params firstName ", firstName: "arvind",
			expRes: []models.Student{
				{ID: 1, FirstName: "arvind", Nationality: "Indian", ContactNumber: 7348761063},
			}},
		{desc: "failure:invalid Query param firstName ", firstName: "123a",
			expErr: errors.New("no rows present in database with this query param")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().GetByFirstName(ctx, tc.firstName).Return(tc.expRes, tc.expErr)

		res, err := mock.Get(ctx, tc.firstName, tc.lastName)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGet_LastName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expRes    []models.Student
		expErr    error
	}{
		{desc: "success:valid Query param  lastName", lastName: "yadav",
			expRes: []models.Student{
				{ID: 1, FirstName: "arvind", LastName: "yadav", Nationality: "Indian", ContactNumber: 7348761063},
			}},
		{desc: "failure:invalid Query param lastName ", lastName: "123a",
			expErr: errors.New("no rows present in database with this query param")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		mockStore.EXPECT().GetByLastName(ctx, tc.lastName).Return(tc.expRes, tc.expErr)

		res, err := mock.Get(ctx, tc.firstName, tc.lastName)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGet_MissingQueryParams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expRes    []models.Student
		expErr    error
	}{
		{desc: "failure:missing Query param firstName and lastName ",
			expErr: errors.New("invalid query params")},
	}

	for i, tc := range testcases {
		ctx := context.Background()
		res, err := mock.Get(ctx, tc.firstName, tc.lastName)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		id        int
		expGetRes models.Student
		expGetErr error
		expErr    error
	}{
		{desc: "success:deleted successfully", id: 1, expGetRes: models.Student{
			ID: 1, FirstName: "arvind", Nationality: "Indian", ContactNumber: 7348761063,
		}},
		{desc: "failure:query error", id: 2, expGetRes: models.Student{
			ID: 2, FirstName: "arvind", Nationality: "Indian", ContactNumber: 7348761063,
		}, expErr: errors.New("query error")},
	}

	for i, tc := range testcases {
		ctx := context.Background()

		mockStore.EXPECT().GetByID(ctx, tc.id).Return(tc.expGetRes, tc.expGetErr)
		mockStore.EXPECT().Delete(ctx, tc.id).Return(tc.expErr)

		err := mock.Delete(ctx, tc.id)

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestDelete_GetErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStudent(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc      string
		id        int
		expGetRes models.Student
		expGetErr error
		expErr    error
	}{
		{desc: "failure:id not present in db result set", id: 1111,
			expGetErr: errors.New("id is not present in db result set"),
			expErr:    errors.New("no sql rows present in db result set")},
	}

	for i, tc := range testcases {
		ctx := context.Background()

		mockStore.EXPECT().GetByID(ctx, tc.id).Return(tc.expGetRes, tc.expGetErr)

		err := mock.Delete(ctx, tc.id)

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}
