package service

import "github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain"

type productServiceImpl struct {
	productRepository domain.ProductRepository
}

func CreateService(r domain.ProductRepository) domain.ProductService {
	return &productServiceImpl{productRepository: r}
}

func (p *productServiceImpl) GetAll() ([]domain.Product, error) {
	products, err := p.productRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p productServiceImpl) GetById(id int64) (*domain.Product, error) {
	product, err := p.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p productServiceImpl) Create(id int64, name string, productType string, description string, quantity int, price float64) (domain.Product, error) {
	newProduct, err := p.productRepository.Create(id, name, productType, description, quantity, price)
	if err != nil {
		return domain.Product{}, nil
	}
	
	return newProduct, nil
}

func (p productServiceImpl) UpdatePrice(id int64, price float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productServiceImpl) Delete(id int64) {
	//TODO implement me
	panic("implement me")
}
