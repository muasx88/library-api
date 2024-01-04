package response

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  StatusResponse `json:"status"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
}

type SuccessResponse struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Token   *string     `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Code    string      `json:"code,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
type StatusResponse struct {
	StatusCode int      `json:"statusCode"`
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Detail     []string `json:"detail"`
}

func OK(c *fiber.Ctx, message string, v interface{}) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{Message: message, Status: fiber.StatusOK, Data: v})
}

func Created(c *fiber.Ctx, message string, v interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{Message: message, Status: fiber.StatusOK, Data: v})
}

func BadRequest(c *fiber.Ctx, msg string, extInfo interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: msg, Status: fiber.StatusBadRequest, Errors: extInfo})
}

func ServerError(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: msg, Status: fiber.StatusInternalServerError})
}

func MsgWithCode(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(ResponseError{Message: msg, Status: code})
}
