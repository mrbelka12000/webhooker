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
	router.HandleFunc("/webhooker", makeWebHookerListHandler(srv)).Methods(http.MethodGet)

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

func makeWebHookerListHandler(srv *service.WebHooker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := srv.List(r.Context(), models.DataListPars{
			Limit:  100,
			Offset: 0,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(body)
	}
}
