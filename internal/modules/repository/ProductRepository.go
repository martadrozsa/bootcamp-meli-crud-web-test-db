package repository

import (
	"context"
	"database/sql"
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
		sqlGetAll,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return products, nil
}

func (r *repositoryImpl) GetById(ctx context.Context, id int64) (*product.Product, error) {
	var prod product.Product

	err := r.conn.GetContext(
		ctx,
		&prod,
		sqlGetById,
		id,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &prod, nil
}

func (r *repositoryImpl) Create(ctx context.Context, name string, productType string, description string, quantity int, price float64) (*product.Product, error) {
	result, err := r.conn.ExecContext(
		ctx,
		sqlCreate,
		name,
		productType,
		description,
		quantity,
		price,
	)

	if err != nil {
		return nil, err
	}

	newProductId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &product.Product{
		Id:          newProductId,
		Name:        name,
		ProductType: productType,
		Description: description,
		Quantity:    quantity,
		Price:       price}, nil
}

func (r *repositoryImpl) UpdatePrice(ctx context.Context, id int64, price float64) error {
	_, err := r.conn.ExecContext(
		ctx,
		sqlUpdatePrice,
		price,
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id int64) error {
	_, err := r.conn.ExecContext(
		ctx,
		sqlDelete,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
