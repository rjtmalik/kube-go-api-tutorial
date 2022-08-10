package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"kube-go-api-tutorial/db/models"
	"net/http"
)

type ProductHandler struct{
	Db *gorm.DB
}

func (p ProductHandler) GetProductRequestHandler(w http.ResponseWriter, r *http.Request){
	var product models.Product
	p.Db.First(&product, "code = ?", "DC")

	json.NewEncoder(w).Encode(map[string]uint{"id": product.ID})
}

func (p ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request){
	p.Db.Create(&models.Product{Code: "DC", Price: 100})
	w.WriteHeader(201)
}