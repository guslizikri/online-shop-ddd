package users

import (
	"context"
	"fmt"
	"log"
	"online-shop-ddd/external/database"
	"online-shop-ddd/infra/response"
	"online-shop-ddd/internal/config"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegisterSuccess(t *testing.T) {
	req := RegisterRequestPayload{
		// ini menggunakan uuid agar emailnya selalu beda
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "abcd1234",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}
func TestRegisterFail(t *testing.T) {
	t.Run("Error email already used", func(t *testing.T) {
		// prepare for duplicate email
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		req := RegisterRequestPayload{
			// ini menggunakan uuid agar emailnya selalu beda
			Email:    email,
			Password: "abcd1234",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)
		// end preparation

		//  sebelumnya regist berhasil dulu,
		// kemudian regist lagi menggunakan email yg sama
		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailExist, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// prepare for login
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		pass := "abcd1234"
		registReq := RegisterRequestPayload{
			Email:    email,
			Password: pass,
		}
		err := svc.register(context.Background(), registReq)
		require.Nil(t, err)
		// end preparation
		loginReq := LoginRequestPayload{
			Email:    email,
			Password: pass,
		}

		token, err := svc.login(context.Background(), loginReq)
		require.Nil(t, err)
		require.NotEmpty(t, token)
		log.Println(token)
	})
}
