package student

import (
	"context"
	"database/sql/driver"
	"errors"
	"log"
	"reflect"
	"testing"

	"student-management-system/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGet(t *testing.T) {
	testcases := []struct {
		desc      string
		expOutput []models.Student
		expRows   *sqlmock.Rows
		expErr    error
	}{
		{desc: "success:get all", expOutput: []models.Student{{ID: 1, FirstName: "arvind", Nationality: "Indian",
			ContactNumber: 7348761063}}, expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow(1, "arvind",
			"", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: nil},
		{desc: "failure:error scanning", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow("abc", "arvind",
			"", "", "", "", "Indian", "", "", "7348761063", "", "", 0), expErr: errors.New("scanning error")},
		{desc: "failure:error select all", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}), expErr: errors.New("error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		mock.ExpectQuery("select * from " + string(models.TableName) + ";").WillReturnRows(tc.expRows).WillReturnError(tc.expErr)

		s := New(db)

		res, err := s.Get(context.TODO())

		if !reflect.DeepEqual(tc.expOutput, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expOutput, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestPost(t *testing.T) {
	testcases := []struct {
		desc    string
		reqData models.Student
		expRes  models.Student
		sqlRes  driver.Result
		expErr  error
	}{
		{desc: "success:posted successfully", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expRes: models.Student{
			ID:            1,
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, sqlRes: sqlmock.NewResult(1, 0), expErr: nil},
		{desc: "failure:query error", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, sqlRes: sqlmock.NewResult(0, 0), expErr: errors.New("query error")},
		{desc: "failure:lastInsertedID error", reqData: models.Student{
			FirstName:     "arvind",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, sqlRes: sqlmock.NewErrorResult(errors.New("lastInsertedId error")), expErr: errors.New("lastInsertedId error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		query := "insert into " + string(models.TableName) + "(first_name,last_name,gender,dob,mother_tongue,nationality,father_name,mother_name,contact_number," +
			"father_occupation,mother_occupation,family_income) values (?,?,?,?,?,?,?,?,?,?,?,?);"
		mock.ExpectExec(query).WithArgs(tc.reqData.FirstName, tc.reqData.LastName, tc.reqData.Gender, tc.reqData.Dob,
			tc.reqData.MotherTongue, tc.reqData.Nationality, tc.reqData.FatherName, tc.reqData.MotherName, tc.reqData.ContactNumber,
			tc.reqData.FatherOccupation, tc.reqData.MotherOccupation, tc.reqData.FamilyIncome).WillReturnResult(tc.sqlRes).
			WillReturnError(tc.expErr)

		s := New(db)

		res, err := s.Post(context.TODO(), &tc.reqData)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGetByFirstAndLastName(t *testing.T) {
	testcases := []struct {
		desc      string
		firstName string
		lastName  string
		expOutput []models.Student
		expRows   *sqlmock.Rows
		expErr    error
	}{
		{desc: "success:get all student with valid first and lastName ", firstName: "arvind", lastName: "yadav", expOutput: []models.Student{{ID: 1,
			FirstName: "arvind", LastName: "yadav", Nationality: "Indian", ContactNumber: 7348761063}},
			expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
				"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
				"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow(1, "arvind",
				"yadav", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: nil},
		{desc: "failure:error scanning", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow("abc", "arvind",
			"", "", "", "", "Indian", "", "", "7348761063", "", "", 0), expErr: errors.New("scanning error")},
		{desc: "failure:error select all", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}), expErr: errors.New("error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		mock.ExpectQuery("select * from"+string(models.TableName)+"where first_name = ? and "+
			"last_name = ?;").WithArgs(tc.firstName, tc.lastName).WillReturnRows(tc.expRows).WillReturnError(tc.expErr)

		s := New(db)

		res, err := s.GetByFirstAndLastName(context.TODO(), tc.firstName, tc.lastName)

		if !reflect.DeepEqual(tc.expOutput, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expOutput, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGetByFirstName(t *testing.T) {
	testcases := []struct {
		desc      string
		firstName string
		expOutput []models.Student
		expRows   *sqlmock.Rows
		expErr    error
	}{
		{desc: "success:get all student with valid firstName ", firstName: "arvind", expOutput: []models.Student{{ID: 1,
			FirstName: "arvind", Nationality: "Indian", ContactNumber: 7348761063}},
			expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
				"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
				"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow(1, "arvind",
				"", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: nil},
		{desc: "failure:error scanning", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow("abc", "arvind",
			"", "", "", "", "Indian", "", "", "7348761063", "", "", 0), expErr: errors.New("scanning error")},
		{desc: "failure:error select all", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}), expErr: errors.New("error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		mock.ExpectQuery("select * from" + string(models.TableName) + "where " +
			"first_name = ?;").WithArgs(tc.firstName).WillReturnRows(tc.expRows).WillReturnError(tc.expErr)

		s := New(db)

		res, err := s.GetByFirstName(context.TODO(), tc.firstName)

		if !reflect.DeepEqual(tc.expOutput, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expOutput, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGetByLastName(t *testing.T) {
	testcases := []struct {
		desc      string
		lastName  string
		expOutput []models.Student
		expRows   *sqlmock.Rows
		expErr    error
	}{
		{desc: "success:get all student with valid firstName ", lastName: "yadav", expOutput: []models.Student{{ID: 1,
			FirstName: "arvind", LastName: "yadav", Nationality: "Indian", ContactNumber: 7348761063}},
			expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
				"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
				"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow(1, "arvind",
				"yadav", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: nil},
		{desc: "failure:error scanning", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow("abc", "arvind",
			"", "", "", "", "Indian", "", "", "7348761063", "", "", 0), expErr: errors.New("scanning error")},
		{desc: "failure:error select all", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}), expErr: errors.New("error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		mock.ExpectQuery("select * from " + string(models.TableName) + "where " +
			"last_name = ?;").WithArgs(tc.lastName).WillReturnRows(tc.expRows).WillReturnError(tc.expErr)

		s := New(db)

		res, err := s.GetByLastName(context.TODO(), tc.lastName)

		if !reflect.DeepEqual(tc.expOutput, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expOutput, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		desc    string
		id      int
		expData models.Student
		expRows *sqlmock.Rows
		expErr  error
	}{
		{desc: "success:valid id", id: 1, expData: models.Student{
			ID:            1,
			FirstName:     "arvind",
			LastName:      "yadav",
			Nationality:   "Indian",
			ContactNumber: 7348761063,
		}, expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
			"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow(1, "arvind",
			"yadav", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: nil},
		{desc: "failure:scanning row error", id: 1,
			expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name",
				"gender", "dob", "mother_tongue", "nationality", "father_name", "mother_name",
				"contact_number", "father_occupation", "mother_occupation", "family_income"}).AddRow("abc", "arvind",
				"yadav", "", "", "", "Indian", "", "", 7348761063, "", "", 0), expErr: errors.New("scanning error")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Println(err.Error())
		}

		ctx := context.TODO()

		s := New(db)

		mock.ExpectQuery("select * from " + string(models.TableName) + "where id = ?;").WithArgs(tc.id).WillReturnRows(tc.expRows).WillReturnError(tc.expErr)

		result, err := s.GetByID(ctx, tc.id)

		if !reflect.DeepEqual(result, tc.expData) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expData, result)
		}

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}
