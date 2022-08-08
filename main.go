package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	srv:= &http.Server{
		Handler: r,
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
