package entity

type User struct {
	Name string
	Password string
	Email string
	Phone string
}
func (a User) GetName( )string {
	return a.Name
}

func (a User) GetPhone() string{
	return a.Phone
}
func (a User) GetEmail() string{
	return a.Email
}
func (a User) GetPassword() string{
	return a.Password
}