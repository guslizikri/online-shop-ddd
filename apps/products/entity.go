package products

import (
	"online-shop-ddd/infra/response"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        int       `db:"id"`
	SKU       string    `db:"sku"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	Stock     int16     `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func NewProductPaginationFromListProductRequest(req ListProductRequestPayload) ProductPagination {
	req = req.GenerateDefaultValue()
	return ProductPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func NewProductFromCreateProductRequest(req CreateProductRequestPayload) Product {
	return Product{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func (p Product) Validate() (err error) {
	err = p.ValidateName()
	if err != nil {
		return
	}
	err = p.ValidatePrice()
	if err != nil {
		return
	}
	err = p.ValidateStock()
	if err != nil {
		return
	}
	return
}
func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		err = response.ErrProductRequired
		return
	}
	if len(p.Name) < 4 {
		err = response.ErrProductInvalid
		return
	}
	return
}
func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		err = response.ErrStockInvalid
		return
	}
	return
}
func (p Product) ValidatePrice() (err error) {
	if p.Price <= 0 {
		err = response.ErrPriceInvalid
		return
	}
	return
}
