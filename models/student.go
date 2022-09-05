package models

type Student struct {
	ID               int    `json:"id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Dob              string `json:"dob,omitempty"`
	MotherTongue     string `json:"mother_tongue,omitempty"`
	Nationality      string `json:"nationality,omitempty"`
	FatherName       string `json:"father_name,omitempty"`
	MotherName       string `json:"mother_name,omitempty"`
	ContactNumber    int    `json:"contact_number,omitempty"`
	FatherOccupation string `json:"father_occupation,omitempty"`
	MotherOccupation string `json:"mother_occupation,omitempty"`
	FamilyIncome     int    `json:"family_income,omitempty"`
}

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

type DbTable string

const (
	TableName DbTable = "student"
)
