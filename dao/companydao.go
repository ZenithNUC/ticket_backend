package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func GetCompanyList() []model.Company {
	var company []model.Company
	db.DB.Find(&company)
	return company
}