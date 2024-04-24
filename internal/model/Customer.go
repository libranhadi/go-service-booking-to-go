package model

type Customer struct {
	Id          uint         `json:"-"`
	Name        string       `json:"name"`
	DateOfBirth string       `json:"tanggal_lahir"`
	PhoneNumber string       `json:"telepon"`
	Nationality string       `json:"kewarganegaraan"`
	Email       string       `json:"email"`
	FamilyList  []FamilyList `json:"keluarga"`
}
