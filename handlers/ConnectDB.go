package handlers

import (
	"blog/models"
	"database/sql"
	"fmt"
)

	func ConnectDB() ([]models.Post, error) {
		dsn := "root:[DB_PASS]@tcp(127.0.0.1:3306)/[DB_NAME]"
        db, err := sql.Open("mysql", dsn)
        if err != nil {
                return nil, fmt.Errorf("failed to open DB: %w", err)
        }
        defer db.Close()

        fmt.Printf("Connected to DB: %T\n", db)
		return GetAllPosts()

	}
