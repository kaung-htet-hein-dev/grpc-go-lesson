package grpc

import (
	"context"
	"product-service/internal/domain"
	"product-service/internal/usecase"
	productpb "product-service/proto"
)

type ProductHandler struct {
	productpb.UnimplementedProductServiceServer
	usecase usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (
	*productpb.GetProductResponse,
	error) {
	product, err := h.usecase.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}

	return &productpb.GetProductResponse{
		Product: &productpb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		},
	}, nil
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (
	*productpb.CreateProductResponse,
	error) {
	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	createdProduct, err := h.usecase.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			Id:          createdProduct.ID,
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       createdProduct.Price,
			Stock:       createdProduct.Stock,
		},
	}, nil
}
