package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain/product"
)

type serviceImpl struct {
	repository product.ProductRepository
}

func CreateProductService(r product.ProductRepository) product.ProductService {
	return &serviceImpl{repository: r}
}

func (s *serviceImpl) GetAll(ctx context.Context) ([]*product.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *serviceImpl) GetById(ctx context.Context, uuid uuid.UUID) (*product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Create(ctx context.Context, uuid uuid.UUID, name string, productType string, description string, quantity int, price float64) (product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) UpdatePrice(ctx context.Context, uuid uuid.UUID, price float64) (product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Delete(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
