
CREATE TABLE IF NOT EXISTS order_items (
    id BIGINT(20) AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT(20) NOT NULL,
    product_id BIGINT(20) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
