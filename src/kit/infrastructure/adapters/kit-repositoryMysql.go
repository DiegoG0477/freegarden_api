package adapters

import (
	database "api-order/src/Database" // Assuming shared DB connection setup
	"api-order/src/kit/domain/entities"
	"database/sql"
	"log" // For logging errors
)

type KitRepositoryMysql struct {
	DB *sql.DB
}

// Reusing the database connection logic
func NewKitRepositoryMysql() (*KitRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &KitRepositoryMysql{DB: db}, nil
}

// Create implements ports.IKit
func (r *KitRepositoryMysql) Create(kit entities.Kit) (entities.Kit, error) {
	query := "INSERT INTO kits (user_id, name, description) VALUES (?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing kit insert statement: %v", err)
		return entities.Kit{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(kit.UserID, kit.Name, kit.Description)
	if err != nil {
		log.Printf("Error executing kit insert statement: %v", err)
		// Consider checking for specific DB errors like foreign key violations if needed
		return entities.Kit{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for kit: %v", err)
		return entities.Kit{}, err
	}

	//kit.ID = int(id) // Set the generated ID

	kit.ID = id

	// We might want to fetch the created_at, but let's keep it simple
	// and rely on the response potentially not needing it immediately,
	// or Get operations fetching it later. The entity has the field if needed.

	return kit, nil
}

// GetByUserID implements ports.IKit
func (r *KitRepositoryMysql) GetByUserID(userID int64) ([]entities.Kit, error) {
	// Select created_at as well, since it's part of the entity
	query := "SELECT kit_id, user_id, name, description, created_at FROM kits WHERE user_id = ?"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error querying kits by user ID %d: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var kits []entities.Kit
	for rows.Next() {
		var kit entities.Kit
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&kit.ID, &kit.UserID, &kit.Name, &kit.Description, &kit.CreatedAt); err != nil {
			log.Printf("Error scanning kit row: %v", err)
			// Decide if one bad row should fail the whole query or just be skipped
			return nil, err // Fail fast for now
		}
		kits = append(kits, kit)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating kit rows: %v", err)
		return nil, err
	}

	// Return empty slice, not nil, if no kits found (standard Go practice)
	if len(kits) == 0 {
		return []entities.Kit{}, nil
	}

	return kits, nil
}
