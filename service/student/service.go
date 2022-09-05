package student

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"student-management-system/models"
	"student-management-system/store"
)

type service struct {
	student store.Student
}

func New(s store.Student) service {
	return service{student: s}
}

func (s service) Post(ctx context.Context, student *models.Student) (models.Student, error) {
	if err := isValidate(student); err != nil {
		return models.Student{}, err
	}

	students, err := s.student.Get(ctx)
	if err != nil {
		return models.Student{}, err
	}

	for i := range students {
		if isDuplicate(&students[i], student) {
			return models.Student{}, errors.New("student already exists")
		}
	}

	return s.student.Post(ctx, student)
}

func (s service) Get(ctx context.Context, firstName, lastName string) ([]models.Student, error) {
	if firstName != "" && lastName != "" {
		students, err := s.student.GetByFirstAndLastName(ctx, firstName, lastName)
		if err != nil {
			return nil, errors.New("no rows present in database with this query params")
		}

		return students, nil
	}

	if firstName != "" {
		students, err := s.student.GetByFirstName(ctx, firstName)
		if err != nil {
			return nil, errors.New("no rows present in database with this query param")
		}

		return students, nil
	}

	if lastName != "" {
		students, err := s.student.GetByLastName(ctx, lastName)
		if err != nil {
			return nil, errors.New("no rows present in database with this query param")
		}

		return students, nil
	}

	return nil, errors.New("invalid query params")
}

func (s service) GetByID(ctx context.Context, id int) (models.Student, error) {
	student, err := s.student.GetByID(ctx, id)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func isDuplicate(s1, s2 *models.Student) bool {
	return s1.FirstName == s2.FirstName && s1.LastName == s2.LastName && s1.Gender == s2.Gender && s1.Dob ==
		s2.Dob && s1.MotherTongue == s2.MotherTongue && s1.Nationality == s2.Nationality && s1.FatherName ==
		s2.FatherName && s1.MotherName == s2.MotherName && s1.ContactNumber == s2.ContactNumber && s1.FatherOccupation ==
		s2.FatherOccupation && s1.MotherOccupation == s2.MotherOccupation && s1.FamilyIncome == s2.FamilyIncome
}

func isValidate(student *models.Student) error {
	switch {
	case !checkMandatoryFields(student.FirstName):
		return errors.New("invalid first name")
	case student.LastName != "" && !checkOptionalFields(student.LastName):
		return errors.New("invalid last name")
	case student.Gender != "" && !checkGender(models.Gender(student.Gender)):
		return errors.New("invalid gender")
	case student.Dob != "" && !checkDob(student.Dob):
		return errors.New("invalid dob")
	case student.MotherTongue != "" && !checkOptionalFields(student.MotherTongue):
		return errors.New("invalid mother tongue")
	case !checkMandatoryFields(student.Nationality):
		return errors.New("invalid nationality")
	case student.FatherName != "" && !checkOptionalFields(student.FatherName):
		return errors.New("invalid father name")
	case student.MotherName != "" && !checkOptionalFields(student.MotherName):
		return errors.New("invalid mother name")
	case !checkContactNumber(student.ContactNumber):
		return errors.New("invalid contact number")
	case student.FatherOccupation != "" && !checkOptionalFields(student.FatherOccupation):
		return errors.New("invalid father occupation")
	case student.MotherOccupation != "" && !checkOptionalFields(student.MotherOccupation):
		return errors.New("invalid mother occupation")
	case student.FamilyIncome != 0 && !checkFamilyIncome(student.FamilyIncome):
		return errors.New("invalid family income")
	default:
		return nil
	}
}

func checkGender(gender models.Gender) bool {
	return gender == models.Male || gender == models.Female || gender == models.Other
}

func checkMandatoryFields(value string) bool {
	if value == "" {
		return false
	}

	for _, v := range value {
		if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			return false
		}
	}

	return true
}

func checkOptionalFields(value string) bool {
	for _, v := range value {
		if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			return false
		}
	}

	return true
}

func checkDob(dob string) bool {
	mm, err := strconv.Atoi(strings.Split(dob, "-")[0])
	if err != nil {
		return false
	}

	if mm < 1 || mm > 12 {
		return false
	}

	dd, err := strconv.Atoi(strings.Split(dob, "-")[1])
	if err != nil {
		return false
	}

	if dd < 1 || dd > 31 {
		return false
	}

	yyyy, err := strconv.Atoi(strings.Split(dob, "-")[2])
	if err != nil {
		return false
	}

	if yyyy < 1000 || yyyy > 9999 {
		return false
	}

	if mm == 4 || mm == 6 || mm == 9 || mm == 11 {
		if dd <= 30 {
			return true
		}

		return false
	}

	return true
}

func checkFamilyIncome(familyIncome int) bool {
	return familyIncome > 0
}

func checkContactNumber(contactNumber int) bool {
	contactNum := strconv.Itoa(contactNumber)

	if len(contactNum) != 10 {
		return false
	}

	return true
}
