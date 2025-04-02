CREATE TABLE IF NOT EXISTS orders  (
    id BIGINT(20) AUTO_INCREMENT PRIMARY KEY,
    client_id BIGINT(20) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'processing', 'completed', 'cancelled') NOT NULL,
    quantity INT NOT NULL,
    product_id BIGINT(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);