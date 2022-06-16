package repository

const (
	sqlGetAll = "SELECT  * FROM product"

	sqlGetById = "SELECT * FROM product WHERE id=?"

	sqlCreate = `
	INSERT INTO 
	product (name, product_type, description, quantity, price) 
	VALUES (?, ?, ?, ?, ?)
	`

	sqlUpdatePrice = `
	UPDATE product 
	SET price=?
	WHERE id=?
	`

	sqlDelete = "DELETE FROM product WHERE id=?"
)
