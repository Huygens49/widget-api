package main

import (
	"fmt"
	"net/http"

	"github.com/Huygens49/widget-api/pkg/reading"
	"github.com/Huygens49/widget-api/pkg/saving"
	"github.com/Huygens49/widget-api/pkg/server/rest"
	"github.com/Huygens49/widget-api/pkg/working"

	"github.com/Huygens49/widget-api/pkg/config"
	"github.com/Huygens49/widget-api/pkg/storage/database"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
TODO
-Error Handling
-Exception Logging
-Maybe a domain object mapper?
-Maybe make a CLI that can migrate the db instead?
*/
func main() {
	fmt.Println("Starting...")

	// Setting up GORM database
	db, err := migrateDatabase()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Setting up dependency injection
	r := database.NewRepository(db)
	l := reading.NewService(r)
	s := saving.NewService(r)
	wrk := working.NewService(r)

	// Setup the router
	router := mux.NewRouter()
	router.StrictSlash(true) // This makes it so "/widgets/" will automatically redirect to "/widgets"

	// widgets
	router.HandleFunc("/widgets", rest.GetWidgets(l)).Methods("GET")
	router.HandleFunc("/widgets", rest.PostWidget(s)).Methods("POST")
	router.HandleFunc("/widgets/{id}", rest.GetWidget(l)).Methods("GET")
	router.HandleFunc("/widgets/{id}", rest.PutWidget(s)).Methods("PUT")

	// work
	router.HandleFunc("/work/{id}", rest.PostWork(wrk)).Methods("POST")

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", router)
}

func migrateDatabase() (*gorm.DB, error) {
	dsn := config.GetDatabaseString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&database.WidgetEntity{})

	if err != nil {
		return nil, err
	}

	return db, nil
}