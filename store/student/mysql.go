package student

import (
	"context"
	"database/sql"

	"student-management-system/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) Get(ctx context.Context) ([]models.Student, error) {
	var students []models.Student

	query := "select * from student;"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Gender, &student.Dob, &student.MotherTongue,
			&student.Nationality, &student.FatherName, &student.MotherName, &student.ContactNumber, &student.FatherOccupation,
			&student.MotherOccupation, &student.FamilyIncome)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s store) GetByID(ctx context.Context, id int) (models.Student, error) {
	var student models.Student

	query := "select * from student where id = ?;"

	err := s.db.QueryRowContext(ctx, query, id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Gender, &student.Dob, &student.MotherTongue,
		&student.Nationality, &student.FatherName, &student.MotherName, &student.ContactNumber, &student.FatherOccupation,
		&student.MotherOccupation, &student.FamilyIncome)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func (s store) GetByFirstAndLastName(ctx context.Context, firstName, lastName string) ([]models.Student, error) {
	var students []models.Student

	query := "select * from student where first_name = ? and last_name = ?;"

	rows, err := s.db.QueryContext(ctx, query, firstName, lastName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Gender, &student.Dob, &student.MotherTongue,
			&student.Nationality, &student.FatherName, &student.MotherName, &student.ContactNumber, &student.FatherOccupation,
			&student.MotherOccupation, &student.FamilyIncome)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s store) GetByFirstName(ctx context.Context, firstName string) ([]models.Student, error) {
	var students []models.Student

	query := "select * from student where first_name = ?;"

	rows, err := s.db.QueryContext(ctx, query, firstName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Gender, &student.Dob, &student.MotherTongue,
			&student.Nationality, &student.FatherName, &student.MotherName, &student.ContactNumber, &student.FatherOccupation,
			&student.MotherOccupation, &student.FamilyIncome)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s store) GetByLastName(ctx context.Context, lastName string) ([]models.Student, error) {
	var students []models.Student

	query := "select * from student where last_name = ?;"

	rows, err := s.db.QueryContext(ctx, query, lastName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Gender, &student.Dob, &student.MotherTongue,
			&student.Nationality, &student.FatherName, &student.MotherName, &student.ContactNumber, &student.FatherOccupation,
			&student.MotherOccupation, &student.FamilyIncome)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s store) Post(ctx context.Context, student *models.Student) (models.Student, error) {
	query := "insert into student (first_name,last_name,gender,dob,mother_tongue,nationality,father_name,mother_name,contact_number," +
		"father_occupation,mother_occupation,family_income) values (?,?,?,?,?,?,?,?,?,?,?,?);"

	res, err := s.db.ExecContext(ctx, query, student.FirstName, student.LastName, student.Gender, student.Dob, student.MotherTongue,
		student.Nationality, student.FatherName, student.MotherName, student.ContactNumber, student.FatherOccupation,
		student.MotherOccupation, student.FamilyIncome)

	if err != nil {
		return models.Student{}, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return models.Student{}, err
	}

	student.ID = int(ID)

	return *student, nil
}
