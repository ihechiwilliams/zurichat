package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"zuri.chat/zccore/data"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/data/write", data.WriteData)
	r.HandleFunc("/data/read", data.ReadData)

	http.Handle("/", r)

	return r
}

func main() {
	fmt.Println("Zuri Chat API running")

	r := Router()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Zuri Chat API - Version 0.0001\n")
}
