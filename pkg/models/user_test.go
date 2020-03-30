package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	t.Run("success create", func(t *testing.T) {
		user, err := CreateUser("andy", "123456", "andy@mail.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "andy", user.Name)
		assert.Equal(t, uint(3), user.ID)
	})

	t.Run("duplicate create", func(t *testing.T) {
		_, err := CreateUser("andy", "123456", "andy@mail.com")
		assert.Error(t, err)
		assert.True(t, IsErrUserAlreadyExist(err))
	})
}
