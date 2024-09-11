package handlers

import (
	"go-hex-mongo/internal/domains/entity"
	"go-hex-mongo/internal/ports"
	"go-hex-mongo/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service ports.IProductService
}

func NewProductHandler(service ports.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(entity.Product)
	if err := c.BodyParser(product); err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}
	if err := h.service.CreateProduct(product); err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	return utils.JSONResponse(c, fiber.StatusCreated, "Berhasil tambah produk", product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(entity.Product)
	if err := c.BodyParser(product); err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}
	if err := h.service.UpdateProduct(id, product); err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	product.ID = id // Ensure ID is included in the response
	return utils.JSONResponse(c, fiber.StatusOK, "Berhasil update produk", product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteProduct(id); err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.GetProductByID(id)
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusNotFound)
	}
	return utils.JSONResponse(c, fiber.StatusOK, "Berhasil mengambil produk berdasarkan ID", product)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	return utils.JSONResponse(c, fiber.StatusOK, "Berhasil mengambil list seluruh produk", products)
}
