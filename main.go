package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kube-go-api-tutorial/db/models"
	"kube-go-api-tutorial/handlers"
	"kube-go-api-tutorial/repository"
	"log"
	"net/http"
	"time"
)

func main(){
	//setup db
	db, err:= gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!=nil{
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})

	//Repository
	productRepo := repository.ProductRepository{Db: db}

	//Handlers
	productHandler := handlers.ProductHandler{DbReader: productRepo, DbWriter: productRepo}

	// setup router
	r := mux.NewRouter()
	r.HandleFunc("/products/1", productHandler.GetProductRequestHandler)
	r.HandleFunc("/products", productHandler.CreateProductHandler)

	srv:= &http.Server{
		Handler: r,
		Addr: ":8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
