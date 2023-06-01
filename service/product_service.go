package service

import "context"

type ProdService struct {
}

func (ps *ProdService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (ps *ProdService) GetProductStock(ctx context.Context, pr *ProductRequest) (*ProductResponse, error) {
	if pr.ProdId == 1 {
		return &ProductResponse{ProdStock: int32(213)}, nil
	}
	return &ProductResponse{
		ProdStock: int32(564),
	}, nil
}
