package repository

import (
	"database/sql"
	"go-service-boooking-to-go/internal/model"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (repo *CustomerRepository) GetCustomer() ([]model.Customer, error) {
	var customers []model.Customer
	rows, err := repo.DB.Query(`
			SELECT
			c.cst_id,
			c.cst_name,
			c.cst_dob,
			c.cst_phoneNum,
			c.cst_email,
			CONCAT (n.nationality_name, ' (', n.nationality_code, ')') as kewarganegaraan
		FROM customers c
		LEFT JOIN nationalities n ON c.nationality_id = n.nationality_id
		WHERE c.deleted_at IS NULL;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.PhoneNumber, &customer.Email, &customer.Nationality)
		if err != nil {
			return nil, err
		}

		familyListRepo := NewFamilyListRepository(repo.DB)
		familyList, err := familyListRepo.GetFamilyByCustomerId(int(customer.Id))
		if err != nil {
			return nil, err
		}
		customer.FamilyList = familyList
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
