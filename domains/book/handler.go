package book

import (
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

func (h handler) GetAllBook(c *fiber.Ctx) error {
	books, err := h.service.GetAll(c.Context())
	if err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	res := NewBooksResponse(books)
	return response.OK(c, "success get books", res)
}

func (h handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, _ := strconv.Atoi(id)

	book, err := h.service.DetailById(c.Context(), intId)
	if err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	res := NewBookResponse(book)
	return response.OK(c, "success get book", res)
}

func (h handler) AddBook(c *fiber.Ctx) error {
	var err error
	var req AddBookRequestPayload

	if err = c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "bad request", err.Error())
	}

	ok, errs := validation.ValidateStruct(req)
	if !ok {
		return response.BadRequest(c, "Bad request", errs)
	}

	err = validation.ValidateISBN(req.ISBN, "isbn")
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}

	if err = h.service.Store(c.Context(), req); err != nil {
		errCode, ok := response.ErrorMap[err]
		if !ok {
			errCode = fiber.StatusInternalServerError
		}

		return response.MsgWithCode(c, errCode, err.Error())
	}

	return response.Created(c, "success add book", nil)
}
