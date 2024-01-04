package auth

import (
	"github.com/muasx88/library-api/internal/response"
	"github.com/muasx88/library-api/internal/validation"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service IService
}

func newHandler(service IService) *handler {
	return &handler{
		service: service,
	}
}

func (h handler) register(c *fiber.Ctx) error {
	var err error
	var req RegisterRequestPayload

	if err = c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "bad request", err.Error())
	}

	ok, errs := validation.ValidateStruct(req)
	if !ok {
		return response.BadRequest(c, "Bad request", errs)
	}

	err = validation.ValidatePassword(req.Password, "password")
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}

	if err = h.service.Register(c.Context(), req); err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	return response.Created(c, "success register", nil)
}

func (h handler) login(c *fiber.Ctx) error {
	var req = LoginRequestPayload{}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}

	token, err := h.service.Login(c.UserContext(), req)
	if err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	return response.OK(c, "login success", map[string]string{
		"access_token": token,
	})
}
