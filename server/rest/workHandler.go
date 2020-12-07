package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Huygens49/widget-api/working"
	"github.com/gorilla/mux"
)

func PostWork(wrk working.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], 10, 0)

		if err != nil {
			fmt.Println("Conversion error")
			return
		}

		wrk.WorkOnWidget(uint(id))

		w.WriteHeader(http.StatusNoContent)
	}
}
