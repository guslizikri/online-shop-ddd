package transactions

import (
	"fmt"
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

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {

	var req = CreateTransactionRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)

	}

	userPublicId := ctx.Locals("PUBLIC_ID")
	req.UserPublicId = fmt.Sprintf("%v", userPublicId)
	if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {
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
		infrafiber.WithMessage("create transactions success"),
	).Send(ctx)
}
func (h handler) GetTransactionByUser(ctx *fiber.Ctx) error {
	userPublicId := fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

	trxs, err := h.svc.TransactionHistories(ctx.UserContext(), userPublicId)
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

	var response = []TransactionHisotryResponse{}

	for _, trx := range trxs {
		response = append(response, trx.ToTransactionHistoryResponse())
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(response),
		infrafiber.WithMessage("get transaction histories success"),
	).Send(ctx)
}
