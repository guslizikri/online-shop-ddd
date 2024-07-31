package transactions

import (
	"context"
	"online-shop-ddd/external/database"
	"online-shop-ddd/internal/config"
	"testing"

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

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "aa4ced3e-aa45-4e83-a2c3-756e806517ed",
			Amount:       2,
			UserPublicId: "08c6e655-5aa7-46ec-84ec-0cd4aa78b033",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})
}
