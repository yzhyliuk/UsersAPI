package models

import (
	"ms/usersAPI/utils"
	"ms/usersAPI/utils/errors"

	"github.com/go-playground/validator/v10"
)

//User : struct that represents User
type User struct {
	ID           int    `json:"id" validate:"omitempty,gt=0"`
	Name         string `json:"name" validate:"omitempty,min=2"`
	LastName     string `json:"lastname" validate:"omitempty,min=2"`
	Email        string `json:"email" gorm:"unique" validate:"omitempty,email"`
	Password     string `json:"password" validate:"omitempty,min=8"`
	CompanyID    int    `json:"companyid" validate:"omitempty,gt=0"`
	DepartmentID int    `json:"departmentid" validate:"omitempty,gt=0"`
	Role         string `json:"role"`
}

//Export convert internal user struct to UserExported omiting privat fields
func (u *User) Export() *UserExported {
	return &UserExported{
		ID:           u.ID,
		Name:         u.Name,
		LastName:     u.LastName,
		Email:        u.Email,
		CompanyID:    u.CompanyID,
		DepartmentID: u.DepartmentID,
		Role:         u.Role,
	}
}

//UserData : extract user main data from User struct
func (u *User) UserData() *UserData {
	return &UserData{
		UserID:       u.ID,
		CompanyID:    u.CompanyID,
		DepartmentID: u.DepartmentID,
		Role:         u.Role,
	}
}

//Merge : merge current User with aother, filling empty fields
func (u *User) Merge(merge *User) {
	utils.Merge(u, merge)
}

//Validate : validates name fieldto be longer/equals 2
func (u *User) Validate() *errors.APIError {
	validator := validator.New()
	err := validator.Struct(u)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}
