package user

type User struct {
	username string
	password string
	email    string
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) GetEmail() string {
	return user.email
}
