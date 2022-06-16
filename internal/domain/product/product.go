package product

import (
	"context"
)

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	ProductType string  `db:"product_type" json:"product_type"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, id int64, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(ctx context.Context, id int64, price float64) (Product, error)
	Delete(ctx context.Context, id int64) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, id int64, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(ctx context.Context, id int64, price float64) (Product, error)
	Delete(ctx context.Context, id int64) error
}
