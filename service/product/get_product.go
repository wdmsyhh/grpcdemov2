package product

import (
	"context"
	"grpcdemov2/proto/product"
)

type ProductService struct {
	product.UnimplementedProductServiceServer
}

func (*ProductService) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	response := &product.GetProductResponse{
		Id:    req.Id,
		Name:  "商品a",
		Price: 1,
	}
	return response, nil
}
