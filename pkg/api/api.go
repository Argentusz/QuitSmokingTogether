package api

import (
	"QuitSmokingTogether/pkg/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type API struct {
	r  *mux.Router
	db storage.DB
}

func New(r *mux.Router, s storage.DB) *API {
	return &API{r: r, db: s}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/users", api.users).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/groups", api.groups).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
