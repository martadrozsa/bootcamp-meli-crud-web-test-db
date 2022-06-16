package product

import (
	"context"
	"github.com/google/uuid"
)

type Product struct {
	UUID        uuid.UUID `db:"uuid" json:"id"`
	Name        string    `json:"name"`
	ProductType string    `db:"product_type" json:"product_type"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, uuid uuid.UUID) (*Product, error)
	Create(ctx context.Context, uuid uuid.UUID, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(ctx context.Context, uuid uuid.UUID, price float64) (Product, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, uuid uuid.UUID) (*Product, error)
	Create(ctx context.Context, uuid uuid.UUID, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(ctx context.Context, uuid uuid.UUID, price float64) (Product, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
}
