package products

import "context"

type Repository interface {
	CreateProduct(ctx context.Context, model Product) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateProduct(ctx context.Context, req ReqCreateProduct) (err error) {
	productEntity := NewProductFromReqCreateProduct(req)

	err = productEntity.Validate()
	if err != nil {
		return
	}

	err = s.repo.CreateProduct(ctx, productEntity)
	if err != nil {
		return
	}
	return
}
