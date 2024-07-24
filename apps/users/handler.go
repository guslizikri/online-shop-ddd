package users

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

func (h handler) register(ctx *fiber.Ctx) (err error) {
	req := RegisterRequestPayload{}
	err = ctx.BodyParser(&req)
	if err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			infrafiber.WithMessage("register fail"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	err = h.svc.register(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			infrafiber.WithMessage("register fail"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("register success"),
	).Send(ctx)
}
func (h handler) login(ctx *fiber.Ctx) (err error) {
	req := LoginRequestPayload{}
	err = ctx.BodyParser(&req)
	if err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			infrafiber.WithMessage("login fail"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	token, err := h.svc.login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			infrafiber.WithMessage("login fail"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
		infrafiber.WithMessage("login success"),
	).Send(ctx)
}
