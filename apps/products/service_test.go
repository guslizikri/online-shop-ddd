package products

import (
	"context"
	"online-shop-ddd/external/database"
	"online-shop-ddd/infra/response"
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

func TestCreateProductSuccess(t *testing.T) {
	req := ReqCreateProduct{
		Name:  "Indomie",
		Price: 3000,
		Stock: 100,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}
func TestCreateProductFail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := ReqCreateProduct{
			Name:  "",
			Price: 3000,
			Stock: 100,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("price must be greater than 0", func(t *testing.T) {
		req := ReqCreateProduct{
			Name:  "indomie",
			Price: 0,
			Stock: 100,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
