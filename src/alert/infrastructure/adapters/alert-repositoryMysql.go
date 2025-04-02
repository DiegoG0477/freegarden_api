package adapters

import (
	database "api-order/src/Database"     // Assuming shared DB connection setup
	"api-order/src/alert/domain/entities" // Adjusted import path
	"database/sql"
	"log"
)

type AlertRepositoryMysql struct {
	DB *sql.DB
}

func NewAlertRepositoryMysql() (*AlertRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &AlertRepositoryMysql{DB: db}, nil
}

// Create implements ports.IAlert
func (r *AlertRepositoryMysql) Create(alert entities.Alert) (entities.Alert, error) {
	query := "INSERT INTO alerts (kit_id, alert_type, message) VALUES (?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing alert insert statement: %v", err)
		return entities.Alert{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(alert.KitID, alert.AlertType, alert.Message)
	if err != nil {
		log.Printf("Error executing alert insert statement: %v", err)
		// Check for specific errors like foreign key violation
		return entities.Alert{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID for alert: %v", err)
		return entities.Alert{}, err
	}

	alert.AlertID = int(id) // Set the generated alert_id

	// Timestamp is not fetched here.

	return alert, nil
}

// GetByKitID implements ports.IAlert
func (r *AlertRepositoryMysql) GetByKitID(kitID int) ([]entities.Alert, error) {
	query := "SELECT alert_id, kit_id, alert_type, message, timestamp FROM alerts WHERE kit_id = ? ORDER BY timestamp DESC" // Order by most recent
	rows, err := r.DB.Query(query, kitID)
	if err != nil {
		// Log specific error for not found if needed, but Query handles it okay
		log.Printf("Error querying alerts by kit ID %d: %v", kitID, err)
		return nil, err // Return error for DB issues
	}
	defer rows.Close()

	var alerts []entities.Alert
	for rows.Next() {
		var alert entities.Alert
		// Ensure Scan order matches SELECT statement
		if err := rows.Scan(&alert.AlertID, &alert.KitID, &alert.AlertType, &alert.Message, &alert.Timestamp); err != nil {
			log.Printf("Error scanning alert row: %v", err)
			return nil, err // Fail fast on scan error
		}
		alerts = append(alerts, alert)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating alert rows: %v", err)
		return nil, err
	}

	// Return empty slice if no alerts found
	if len(alerts) == 0 {
		return []entities.Alert{}, nil
	}

	return alerts, nil
}
