package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain/product"
)

type repositoryImpl struct {
	conn *sqlx.DB
}

func CreateProductRepository(conn *sqlx.DB) product.ProductRepository {
	return &repositoryImpl{conn}
}

func (r *repositoryImpl) GetAll(ctx context.Context) ([]*product.Product, error) {
	var products []*product.Product

	err := r.conn.SelectContext(
		ctx,
		&products,
		sqlFindAll,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return products, nil
}

func (r *repositoryImpl) GetById(ctx context.Context, uuid uuid.UUID) (*product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Create(ctx context.Context, uuid uuid.UUID, name string, productType string, description string, quantity int, price float64) (product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) UpdatePrice(ctx context.Context, uuid uuid.UUID, price float64) (product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Delete(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
