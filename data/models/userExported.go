package models

//UserExported : struct that represents User only for internal use (inside microservice)
type UserExported struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	CompanyID int    `json:"companyid"`
}
