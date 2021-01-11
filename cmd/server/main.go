package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Huygens49/widget-api/pkg/config"
	"github.com/Huygens49/widget-api/pkg/read"
	"github.com/Huygens49/widget-api/pkg/server/rest"
	"github.com/Huygens49/widget-api/pkg/storage/database"
	"github.com/Huygens49/widget-api/pkg/working"
	"github.com/Huygens49/widget-api/pkg/write"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	// Setting up SQL database
	db, err := sql.Open("postgres", config.GetDatabaseString())
	defer db.Close()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Setting up dependency injection
	r := database.NewRepository(db)
	l := read.NewService(r)
	s := write.NewService(r)
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
