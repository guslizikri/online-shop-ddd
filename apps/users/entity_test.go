package users

import (
	"log"
	"online-shop-ddd/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestValidateUserEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "my@email.com",
			Password: "mysecretpass",
		}
		err := userEntity.Validate()
		require.Nil(t, err)
	})

	// test negative case email
	t.Run("email required", func(t *testing.T) {
		userEntity := UserEntity{
			Password: "mysecretpass",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})
	t.Run("invalid email", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "myemail.com",
			Password: "mysecretpass",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	// test negative case password
	t.Run("password required", func(t *testing.T) {
		userEntity := UserEntity{
			Email: "my@email.com",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})
	t.Run("invalid pass lenght", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "my@email.com",
			Password: "pass",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})
}

func TestEncryptPass(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "my@email.com",
			Password: "pass",
		}
		err := userEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", userEntity)
	})
}
