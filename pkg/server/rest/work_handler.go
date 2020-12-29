package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Huygens49/widget-api/pkg/working"
	"github.com/gorilla/mux"
)

func PostWork(wrk working.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			fmt.Println("Conversion error")
			return
		}

		wrk.WorkOnWidget(uint(id))

		w.WriteHeader(http.StatusNoContent)
	}
}
