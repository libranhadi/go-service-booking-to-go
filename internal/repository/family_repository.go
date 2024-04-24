package repository

import (
	"database/sql"
	"go-service-boooking-to-go/internal/model"
)

type FamilyListRepository struct {
	DB *sql.DB
}

func NewFamilyListRepository(db *sql.DB) *FamilyListRepository {
	return &FamilyListRepository{DB: db}
}

func (repo FamilyListRepository) GetFamilyByCustomerId(id int) ([]model.FamilyList, error) {
	var familyList []model.FamilyList

	rows, err := repo.DB.Query("SELECT fl_name, fl_relation, fl_dob FROM family_lists WHERE cst_id = ? AND deleted_at IS NULL", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var family model.FamilyList
		err := rows.Scan(&family.Name, &family.Relation, &family.DateOfBirth)
		if err != nil {
			return nil, err
		}
		familyList = append(familyList, family)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return familyList, nil
}
