package student

import (
	"context"
	"errors"

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
	if err := IsValidate(student); err != nil {
		return models.Student{}, err
	}

	students, err := s.student.Get(ctx)
	if err != nil {
		return models.Student{}, err
	}

	for i := range students {
		if IsDuplicate(students[i], student) {
			return models.Student{}, errors.New("student already exists")
		}
	}

	return s.student.Post(ctx, student)
}

func IsDuplicate(s1, s2 models.Student) bool {
	return s1.FirstName == s2.FirstName && s1.LastName == s2.LastName && s1.Gender == s2.Gender && s1.Dob ==
		s2.Dob && s1.MotherTongue == s2.MotherTongue && s1.Nationality == s2.Nationality && s1.FatherName ==
		s2.FatherName && s1.MotherName == s2.MotherName && s1.ContactNumber == s2.ContactNumber && s1.FatherOccupation ==
		s2.FatherOccupation && s1.MotherOccupation == s2.MotherOccupation && s1.FamilyIncome == s2.FamilyIncome
}

func IsValidate(student models.Student) error {
	switch {
	case checkFirstName(student.FirstName):
		return errors.New("invalid first name")
	case checkLastName(student.LastName):
		return errors.New("invalid last name")
	case checkGender(student.Gender):
		return errors.New("invalid gender")
	case checkDob(student.Dob):
		return errors.New("invalid dob")
	case checkMotherTongue(student.MotherTongue):
		return errors.New("invalid mother tongue")
	case checkNationality(student.Nationality):
		return errors.New("invalid nationality")
	case checkFatherName(student.FatherName):
		return errors.New("invalid father name")
	case checkMotherName(student.MotherName):
		return errors.New("invalid mother name")
	case checkContactNumber(student.ContactNumber):
		return errors.New("invalid contact number")
	case checkFatherOccupation(student.FatherOccupation):
		return errors.New("invalid father occupation")
	case checkMotherOccupation(student.MotherOccupation):
		return errors.New("invalid mother occupation")
	case checkFamilyIncome(student.FamilyIncome):
		return errors.New("invalid family income")
	default:
		return nil
	}
}
