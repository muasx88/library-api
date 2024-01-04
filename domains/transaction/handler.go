package transaction

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/muasx88/library-api/internal/response"
	"github.com/muasx88/library-api/internal/validation"
)

type handler struct {
	service IService
}

func newHandler(service IService) *handler {
	return &handler{
		service: service,
	}
}

func (h handler) GetUserHistoryTransaction(c *fiber.Ctx) error {
	id := c.Params("user_id")
	intId, _ := strconv.Atoi(id)

	histories, err := h.service.UserHistoryTransaction(c.Context(), intId)
	if err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	res := NewUserHistoryTransactionResponse(histories)
	return response.OK(c, "success get histories", res)
}

func (h handler) BorrowBook(c *fiber.Ctx) error {
	var err error
	var req BorrowRequest
	var userIdStr string

	if err = c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "bad request", err.Error())
	}

	ok, errs := validation.ValidateStruct(req)
	if !ok {
		return response.BadRequest(c, "Bad request", errs)
	}

	userIdStr = fmt.Sprintf("%v", c.Locals("id"))
	userIdInt, _ := strconv.Atoi(userIdStr)
	req.UserId = userIdInt

	if err = h.service.Borrow(c.Context(), req); err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	return response.OK(c, "success borrow book", nil)
}

func (h handler) ReturnBook(c *fiber.Ctx) error {
	var err error
	var req ReturnRequest
	var userIdStr string

	if err = c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "bad request", err.Error())
	}

	ok, errs := validation.ValidateStruct(req)
	if !ok {
		return response.BadRequest(c, "Bad request", errs)
	}

	userIdStr = fmt.Sprintf("%v", c.Locals("id"))
	userIdInt, _ := strconv.Atoi(userIdStr)
	req.UserId = userIdInt

	if err = h.service.Return(c.Context(), req); err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	return response.OK(c, "success return book", nil)
}
