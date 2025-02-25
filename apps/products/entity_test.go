package products

import (
	"online-shop-ddd/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		product := Product{
			Name:  "Indomie",
			Price: 3000,
			Stock: 100,
		}
		err := product.Validate()

		require.Nil(t, err)

	})
	t.Run("name product required", func(t *testing.T) {
		product := Product{
			Name:  "",
			Price: 3000,
			Stock: 100,
		}
		err := product.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)

	})
	t.Run("product must minimum 4 char", func(t *testing.T) {
		product := Product{
			Name:  "Ind",
			Price: 3000,
			Stock: 100,
		}
		err := product.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)

	})
	t.Run("price invalid", func(t *testing.T) {
		product := Product{
			Name:  "Indomie",
			Price: 0,
			Stock: 100,
		}
		err := product.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)

	})
	t.Run("stock invalid", func(t *testing.T) {
		product := Product{
			Name:  "Indomie",
			Price: 199,
			Stock: 0,
		}
		err := product.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)

	})
}
