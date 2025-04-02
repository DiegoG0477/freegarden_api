package adapters

import (
	"database/sql"
	"errors"
	"time"

	database "api-order/src/Database"
	"api-order/src/order/domain/entities"
)

type OrderRepositoryMysql struct {
	db *sql.DB
}

func NewOrderRepositoryMysql() (*OrderRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &OrderRepositoryMysql{db: db}, nil
}

// Crear una nueva orden con un solo producto
func (r *OrderRepositoryMysql) Create(order entities.Order) (entities.Order, error) {
	query := `INSERT INTO orders (client_id, product_id, quantity, total_price, status) 
              VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, order.Client_id, order.Product_id, order.Quantity, order.Total_price, order.Status)
	if err != nil {
		return entities.Order{}, err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return entities.Order{}, err
	}

	order.ID = int(id)
	return order, nil
}

// Actualizar el estado de un pedido
func (r *OrderRepositoryMysql) UpdateStatus(id int64, status string) (entities.Order, error) {
	query := `UPDATE orders SET status = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, time.Now(), id)
	if err != nil {
		return entities.Order{}, err
	}

	return r.GetById(id)
}

// Obtener una orden por su ID
func (r *OrderRepositoryMysql) GetById(id int64) (entities.Order, error) {
	query := `SELECT id, client_id, product_id, quantity, total_price, status FROM orders WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var order entities.Order
	err := row.Scan(&order.ID, &order.Client_id, &order.Product_id, &order.Quantity, &order.Total_price, &order.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Order{}, errors.New("order not found")
		}
		return entities.Order{}, err
	}

	return order, nil
}

// Listar pedidos por cliente
func (r *OrderRepositoryMysql) ListOrdersByClient(client_id int64) ([]entities.Order, error) {
	query := `SELECT id, client_id, product_id, quantity, total_price, status FROM orders WHERE client_id = ?`
	rows, err := r.db.Query(query, client_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.ID, &order.Client_id, &order.Product_id, &order.Quantity, &order.Total_price, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// Eliminar un pedido por ID
func (r *OrderRepositoryMysql) Delete(id int64) (bool, error) {
	query := `DELETE FROM orders WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
