package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/webhooker/internal/service"
)

func RegisterRoutes(srv *service.WebHooker) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/webhooker", makeWebHookerCreateHandler(srv)).Methods(http.MethodPost)

	return router
}

func makeWebHookerCreateHandler(srv *service.WebHooker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("ok"))
	}
}
