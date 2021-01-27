package models

import (
	"ms/usersAPI/utils/errors"
	"regexp"
)

//User : struct that represents User
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CompanyID int    `json:"companyid"`
}

//Export convert internal user struct to UserExported omiting privat fields
func (u *User) Export() *UserExported {
	return &UserExported{
		ID:        u.ID,
		Name:      u.Name,
		LastName:  u.LastName,
		Email:     u.Email,
		CompanyID: u.CompanyID,
	}
}

//Merge : merge current User with aother, filling empty fields
func (u *User) Merge(merge *User) {
	if u.Name == "" {
		u.Name = merge.Name
	}
	if u.LastName == "" {
		u.LastName = merge.LastName
	}
	if u.Email == "" {
		u.Email = merge.Email
	}
	if u.Password == "" {
		u.Password = merge.Password
	}
	if u.CompanyID == 0 {
		u.CompanyID = merge.CompanyID
	}
}

//Validate : validates name fieldto be longer/equals 2
func (u *User) Validate() *errors.APIError {
	if len(u.Name) < 2 {
		return errors.NewBadRequestError("User validation error: name field should be at least 2 characters long")
	}
	if len(u.LastName) < 2 {
		return errors.NewBadRequestError("User validation error: last name field should be at least 2 characters long")
	}
	if u.CompanyID <= 0 {
		return errors.NewBadRequestError("User validation error: 'companyid' filed should be positive integer")
	}
	if len(u.Password) < 10 {
		return errors.NewBadRequestError("User validation error: 'password' filed should at least 10 characters long")
	}
	return u.validateEmail()
}

func (u *User) validateEmail() *errors.APIError {
	regularExp := regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	if regularExp.MatchString(u.Email) {
		return nil
	}
	return errors.NewBadRequestError("User validation error: 'email' field is not valid")
}
