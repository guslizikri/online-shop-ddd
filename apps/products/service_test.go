package products

import (
	"context"
	"log"
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
	req := CreateProductRequestPayload{
		Name:  "Indomie",
		Price: 3000,
		Stock: 100,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}
func TestCreateProductFail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Price: 3000,
			Stock: 100,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("price must be greater than 0", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "indomie",
			Price: 0,
			Stock: 100,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}

func TestListProductSuccess(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestProductDetailSuccess(t *testing.T) {
	// prepare to get sku for product detail
	req := CreateProductRequestPayload{
		Name:  "Baju Baru",
		Stock: 10,
		Price: 10_000,
	}

	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)
	// end prepare

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)

	log.Printf("%+v", product)
}
