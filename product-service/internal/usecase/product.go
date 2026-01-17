package usecase

import "product-service/internal/domain"

type ProductUsecase interface {
	GetProduct(id int64) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}

type productUsecase struct{}

func NewProductUsecase() ProductUsecase {
	return &productUsecase{}
}

func (u *productUsecase) GetProduct(id int64) (*domain.Product, error) {
	return &domain.Product{
		ID: id,
	}, nil
}

func (u *productUsecase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	product.ID = 1
	return product, nil
}
