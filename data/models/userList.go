package models

//UserList : List of users
type UserList []*User

//Export : converts an UserList to UserExportedList
func (u UserList) Export() UserExportedList {
	userExportedList := make(UserExportedList, 0)
	for _, value := range u {
		userExportedList = append(userExportedList, value.Export())
	}
	return userExportedList
}
