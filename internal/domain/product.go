package domain

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	ProductType string  `json:"product_type"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(int64) (*Product, error)
	Create(id int64, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(id int64, price float64) (Product, error)
	Delete(id int64)
}

type ProductService interface {
	GetAll() ([]Product, error)
	GetById(int64) (*Product, error)
	Create(id int64, name string, productType string, description string, quantity int, price float64) (Product, error)
	UpdatePrice(id int64, price float64) (Product, error)
	Delete(id int64)
}
