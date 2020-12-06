package main

import (
	"fmt"

	"github.com/Huygens49/widget-api/config"
	"github.com/Huygens49/widget-api/listing"
	"github.com/Huygens49/widget-api/storage/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting...")

	db, err := migrateDatabase()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	repository := database.NewRepository(db)

	// Test query
	service := listing.NewService(repository)
	ws, err := service.GetAllWidgets()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for _, w := range ws {
		fmt.Printf("Result: %d, %s, %s\n", w.ID, w.Description, w.Owner)
	}

	fmt.Println("Done!")
}

func migrateDatabase() (*gorm.DB, error) {
	dsn := config.GetDatabaseString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// TODO real logging?
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&database.WidgetEntity{})

	// TODO real logging?
	if err != nil {
		return nil, err
	}

	return db, nil
}
