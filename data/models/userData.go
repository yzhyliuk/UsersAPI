package models

//UserData : data to put in Access Token
type UserData struct {
	UserID       int    `json:"id"`
	CompanyID    int    `json:"companyid"`
	DepartmentID int    `json:"departmentid"`
	Role         string `json:"role"`
}
