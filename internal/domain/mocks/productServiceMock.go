package mocks

import (
	"context"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain/product"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (p *ProductServiceMock) GetAll(ctx context.Context) ([]*product.Product, error) {
	args := p.Called(ctx)

	var prod []*product.Product
	rf, ok := args.Get(0).(func(ctx context.Context) []*product.Product)

	if ok {
		prod = rf(ctx)
	} else {
		if args.Get(0) != nil {
			prod = args.Get(0).([]*product.Product)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(context.Context) error); ok {
		err = rf(ctx)
	} else {
		err = args.Error(1)
	}

	return prod, err
}

func (p *ProductServiceMock) GetById(ctx context.Context, id int64) (*product.Product, error) {
	args := p.Called(ctx, id)

	var prod *product.Product
	rf, ok := args.Get(0).(func(context.Context, int64) *product.Product)

	if ok {
		prod = rf(ctx, id)
	} else {
		if args.Get(0) != nil {
			prod = args.Get(0).(*product.Product)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(context.Context, int64) error); ok {
		err = rf(ctx, id)
	} else {
		err = args.Error(1)
	}

	return prod, err
}

func (p *ProductServiceMock) Create(ctx context.Context, name string, productType string, description string, quantity int, price float64) (*product.Product, error) {
	args := p.Called(ctx, name, productType, description, quantity, price)

	var prod *product.Product
	rf, ok := args.Get(0).(func(ctx context.Context, name string, productType string, description string, quantity int, price float64) *product.Product)

	if ok {
		prod = rf(ctx, name, productType, description, quantity, price)
	} else {
		if args.Get(0) != nil {
			prod = args.Get(0).(*product.Product)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(context.Context, string, string, string, int, float64) error); ok {
		err = rf(ctx, name, productType, description, quantity, price)
	} else {
		err = args.Error(1)
	}

	return prod, err
}

func (p *ProductServiceMock) UpdatePrice(ctx context.Context, id int64, price float64) error {
	args := p.Called(ctx, id)

	var err error
	rf, ok := args.Get(0).(func(context.Context, int64, float64) error)

	if ok {
		err = rf(ctx, id, price)
	} else {
		err = args.Error(0)
	}

	return err
}

func (p *ProductServiceMock) Delete(ctx context.Context, id int64) error {
	args := p.Called(ctx, id)

	var err error
	rf, ok := args.Get(0).(func(context.Context, int64) error)

	if ok {
		err = rf(ctx, id)
	} else {
		err = args.Error(0)
	}

	return err
}
