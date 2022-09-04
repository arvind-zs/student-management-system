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

func (s service) Post(ctx context.Context, student models.Student) (models.Student, error) {
	if err := isValidate(student); err != nil {
		return models.Student{}, err
	}

	students, err := s.student.Get(ctx)
	if err != nil {
		return models.Student{}, err
	}

	for i := range students {
		if isDuplicate(students[i], student) {
			return models.Student{}, errors.New("student already exists")
		}
	}

	return s.student.Post(ctx, student)
}

func isDuplicate(s1, s2 models.Student) bool {
	return s1.FirstName == s2.FirstName && s1.LastName == s2.LastName && s1.Gender == s2.Gender && s1.Dob ==
		s2.Dob && s1.MotherTongue == s2.MotherTongue && s1.Nationality == s2.Nationality && s1.FatherName ==
		s2.FatherName && s1.MotherName == s2.MotherName && s1.ContactNumber == s2.ContactNumber && s1.FatherOccupation ==
		s2.FatherOccupation && s1.MotherOccupation == s2.MotherOccupation && s1.FamilyIncome == s2.FamilyIncome
}

func isValidate(student models.Student) error {
	switch {
	case !checkFirstName(student.FirstName):
		return errors.New("invalid first name")
	case student.LastName != "" && !checkLastName(student.LastName):
		return errors.New("invalid last name")
	case student.Gender != "" && !checkGender(student.Gender):
		return errors.New("invalid gender")
	case student.Dob != "" && !checkDob(student.Dob):
		return errors.New("invalid dob")
	case student.MotherTongue != "" && !checkMotherTongue(student.MotherTongue):
		return errors.New("invalid mother tongue")
	case !checkNationality(student.Nationality):
		return errors.New("invalid nationality")
	case student.FatherName != "" && !checkFatherName(student.FatherName):
		return errors.New("invalid father name")
	case student.MotherName != "" && !checkMotherName(student.MotherName):
		return errors.New("invalid mother name")
	case !checkContactNumber(student.ContactNumber):
		return errors.New("invalid contact number")
	case student.FatherOccupation != "" && !checkFatherOccupation(student.FatherOccupation):
		return errors.New("invalid father occupation")
	case student.MotherOccupation != "" && !checkMotherOccupation(student.MotherOccupation):
		return errors.New("invalid mother occupation")
	case student.FamilyIncome != 0 && !checkFamilyIncome(student.FamilyIncome):
		return errors.New("invalid family income")
	default:
		return nil
	}
}

func checkGender(gender string) bool {
	return gender == "M" || gender == "F" || gender == "O"
}

func checkFirstName(firstName string) bool {
	if firstName == "" {
		return false
	}

	for _, value := range firstName {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkNationality(nationality string) bool {
	if nationality == "" {
		return false
	}

	for _, value := range nationality {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkLastName(lastName string) bool {
	for _, value := range lastName {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
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

func checkMotherTongue(motherTongue string) bool {
	for _, value := range motherTongue {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkFatherName(fatherName string) bool {
	for _, value := range fatherName {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkMotherName(motherName string) bool {
	for _, value := range motherName {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkFatherOccupation(fatherOccupation string) bool {
	for _, value := range fatherOccupation {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
	}

	return true
}

func checkMotherOccupation(motherOccupation string) bool {
	for _, value := range motherOccupation {
		if !((value >= 65 && value <= 90) || (value >= 97 && value <= 122)) {
			return false
		}
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
