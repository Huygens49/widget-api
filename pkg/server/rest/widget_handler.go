package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Huygens49/widget-api/pkg/read"
	"github.com/Huygens49/widget-api/pkg/write"
	"github.com/gorilla/mux"
)

func GetWidgets(rs read.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		widgets, err := rs.GetAllWidgets()

		if err != nil {
			fmt.Println(err)
			return
		}

		json.NewEncoder(w).Encode(widgets)
	}
}

func GetWidget(rs read.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			fmt.Println("Conversion error")
			return
		}

		widget, err := rs.GetWidget(uint(id))

		if err != nil {
			fmt.Println(err)
			return
		}

		json.NewEncoder(w).Encode(widget)
	}
}

func PostWidget(s write.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var widget write.Widget
		err := json.NewDecoder(r.Body).Decode(&widget)

		if err != nil {
			fmt.Println(err)
			return
		}

		newWidget, err := s.AddWidget(widget)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newWidget)
	}
}

func PutWidget(s write.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var widget write.Widget
		err := json.NewDecoder(r.Body).Decode(&widget)

		if err != nil {
			fmt.Println(err)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			fmt.Println("Conversion error")
			return
		}

		s.UpdateWidget(uint(id), widget)

		w.WriteHeader(http.StatusNoContent)
	}
}
