package service

import (
	"context"
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

func (s *serviceImpl) GetById(ctx context.Context, id int64) (*product.Product, error) {
	album, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *serviceImpl) Create(ctx context.Context, name string, productType string, description string, quantity int, price float64) (*product.Product, error) {
	newProduct, err := s.repository.Create(ctx, name, productType, description, quantity, price)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (s *serviceImpl) UpdatePrice(ctx context.Context, id int64, price float64) error {
	return s.repository.UpdatePrice(ctx, id, price)
}

func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
