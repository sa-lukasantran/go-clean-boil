package main

import (
	"go-basesource/src/adapters"
	"go-basesource/src/entity"
	"log"
)

func main() {
	// Initialize storage
	adapters.StorageInit()

	// Get database instance
	db := adapters.GetDatabase()

	// Run migrations
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&entity.Example{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
