package services

import (
	"go-hex-mongo/internal/domains/entity"
	"go-hex-mongo/internal/ports"
)

// ProductServiceImpl is the concrete implementation of the ProductService interface
type ProductServiceImpl struct {
	repository ports.IProductRepo
}

// NewProductService creates a new ProductServiceImpl
func NewProductServiceImpl(repo ports.IProductRepo) ports.IProductService {
	return &ProductServiceImpl{repository: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *entity.Product) error {
	return s.repository.CreateProduct(product)
}

func (s *ProductServiceImpl) UpdateProduct(id string, product *entity.Product) error {
	return s.repository.UpdateProduct(id, product)
}

func (s *ProductServiceImpl) DeleteProduct(id string) error {
	return s.repository.DeleteProduct(id)
}

func (s *ProductServiceImpl) GetProductByID(id string) (*entity.Product, error) {
	return s.repository.GetProductByID(id)
}

func (s *ProductServiceImpl) GetAllProducts() ([]*entity.Product, error) {
	return s.repository.GetAllProducts()
}
