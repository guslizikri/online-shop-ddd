package products

import (
	"net/http"
	infrafiber "online-shop-ddd/infra/fiber"
	"online-shop-ddd/infra/response"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	req := CreateProductRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err = h.svc.CreateProduct(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create product success"),
	).Send(ctx)
}
func (h handler) GetListProduct(ctx *fiber.Ctx) error {
	req := ListProductRequestPayload{}

	err := ctx.QueryParser(&req)
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	productListResponse := NewProductListResponseFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get list product success"),
		infrafiber.WithPayload(productListResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetDetailProduct(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.ProductDetail(ctx.UserContext(), sku)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	productDetail := ProductDetailResponse{
		Id:        product.Id,
		Name:      product.Name,
		SKU:       product.SKU,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get detail product success"),
		infrafiber.WithPayload(productDetail),
	).Send(ctx)
}
