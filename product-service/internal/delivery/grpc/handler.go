package grpc

import (
	"context"
	"product-service/internal/domain"
	"product-service/internal/usecase"
	productpb "product-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	product, err := h.usecase.GetProduct(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "product not found")
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
	if req.Price <= 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"price must be greater than zero",
		)
	}

	if req.Stock < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"stock cannot be negative",
		)
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	createdProduct, err := h.usecase.CreateProduct(product)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"failed to create product",
		)
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
