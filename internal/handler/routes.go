package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/webhooker/internal/models"
	"github.com/mrbelka12000/webhooker/internal/service"
)

func RegisterRoutes(srv *service.WebHooker) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/webhooker", makeWebHookerCreateHandler(srv)).Methods(http.MethodPost)

	return router
}

func makeWebHookerCreateHandler(srv *service.WebHooker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.Data

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = srv.Create(r.Context(), &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
