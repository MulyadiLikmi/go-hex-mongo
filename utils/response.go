package utils

import "github.com/gofiber/fiber/v2"

// JSONResponse sends a JSON response with status code, message, and optional data
func JSONResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

// HandleError handles errors and sends a standardized error response
func HandleError(c *fiber.Ctx, err error, statusCode int) error {
	return JSONResponse(c, statusCode, err.Error(), nil)
}
