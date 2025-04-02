package adapters

import (
	"database/sql"
	database "api-order/src/Database"
	"api-order/src/product/domain/entities"
)

type ProductRepositoryMysql struct {
	db *sql.DB
}

func NewProductRepositoryMysql() (*ProductRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &ProductRepositoryMysql{db: db}, nil
}

func (r *ProductRepositoryMysql) Create(product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return entities.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.Product{}, err
	}

	product.ID = int(id)
	return product, nil
}

func (r *ProductRepositoryMysql) GetAll() ([]entities.Product, error) {
	query := `SELECT id, name, description, price, stock FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
