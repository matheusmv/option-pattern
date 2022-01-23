package main

import (
	"fmt"

	"github.com/matheusmv/design-patterns/user"
)

func main() {
	user := user.NewUser(
		user.WithUsername("username"),
		user.WithPassword("password"),
		user.WithEmail("user@email.com"),
	)

	fmt.Println(user)
}
