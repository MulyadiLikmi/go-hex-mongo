package ports

import "go-hex-mongo/internal/domains/entity"

type IProductService interface {
	CreateProduct(product *entity.Product) error
	UpdateProduct(id string, product *entity.Product) error
	DeleteProduct(id string) error
	GetProductByID(id string) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
}
