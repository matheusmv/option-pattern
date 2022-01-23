package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("when no option is provided", func(t *testing.T) {
		expectedUser := new(User)
		expectedUser.username = defaultUsername
		expectedUser.password = defaultPassword
		expectedUser.email = defatulEmail

		t.Run("should return a User with default parameters", func(t *testing.T) {
			user := NewUser()

			assert.Equal(t, user, expectedUser)
			assert.Equal(t, user.GetUsername(), defaultUsername)
			assert.Equal(t, user.GetPassword(), defaultPassword)
			assert.Equal(t, user.GetEmail(), defatulEmail)
		})
	})
}

func TestWithUsername(t *testing.T) {
	providedUsername := "Jhon"

	t.Run("when username option is provided", func(t *testing.T) {
		expectedUser := new(User)
		expectedUser.username = providedUsername
		expectedUser.password = defaultPassword
		expectedUser.email = defatulEmail

		t.Run("should return a User with the appropriate options defined", func(t *testing.T) {
			user := NewUser(
				WithUsername(providedUsername),
			)

			assert.Equal(t, user, expectedUser)
			assert.Equal(t, user.GetUsername(), providedUsername)
		})
	})
}

func TestWithPassword(t *testing.T) {
	providedPassword := "123456"

	t.Run("when password option is provided", func(t *testing.T) {
		expectedUser := new(User)
		expectedUser.username = defaultUsername
		expectedUser.password = providedPassword
		expectedUser.email = defatulEmail

		t.Run("should return a User with the appropriate options defined", func(t *testing.T) {
			user := NewUser(
				WithPassword(providedPassword),
			)

			assert.Equal(t, user, expectedUser)
			assert.Equal(t, user.GetPassword(), providedPassword)
		})
	})
}

func TestWithEmail(t *testing.T) {
	providedEmail := "jhon@email.com"

	t.Run("when email option is provided", func(t *testing.T) {
		expectedUser := new(User)
		expectedUser.username = defaultUsername
		expectedUser.password = defaultPassword
		expectedUser.email = providedEmail

		t.Run("should return a User with the appropriate options defined", func(t *testing.T) {
			user := NewUser(
				WithEmail(providedEmail),
			)

			assert.Equal(t, user, expectedUser)
			assert.Equal(t, user.GetEmail(), providedEmail)
		})
	})
}
