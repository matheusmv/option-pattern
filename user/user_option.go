package user

type UserOption func(*User)

const (
	defaultUsername = "default username"
	defaultPassword = "default password"
	defaultEmail    = "default email"
)

func NewUser(opts ...UserOption) *User {
	user := &User{
		username: defaultUsername,
		password: defaultPassword,
		email:    defaultEmail,
	}

	for _, opt := range opts {
		opt(user)
	}

	return user
}

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.username = username
	}
}

func WithPassword(password string) UserOption {
	return func(u *User) {
		u.password = password
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.email = email
	}
}
