package repository

import "github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain"

type productRepositoryImpl struct {
}

func CreateRepository() domain.ProductRepository {
	return &productRepositoryImpl{}
}

func (p productRepositoryImpl) GetAll() ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepositoryImpl) GetById(id int64) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepositoryImpl) Create(id int64, name string, productType string, description string, quantity int, price float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepositoryImpl) UpdatePrice(id int64, price float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepositoryImpl) Delete(id int64) {
	//TODO implement me
	panic("implement me")
}
