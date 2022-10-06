package main

import (
	"QuitSmokingTogether/pkg/api"
	"QuitSmokingTogether/pkg/storage/pgsql"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	s, err := pgsql.NewStorage(CONN)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := api.New(mux.NewRouter(), s)
	a.Handle()
	a.ListenAndServe(ADDR)
}
