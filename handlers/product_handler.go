package handlers

import (
	"encoding/json"
	"kube-go-api-tutorial/db/models"
	"kube-go-api-tutorial/repository_interfaces"
	"net/http"
)

type ProductHandler struct{
	DbReader repository_interfaces.ProductReader
	DbWriter repository_interfaces.ProductWriter
}

func (p ProductHandler) GetProductRequestHandler(w http.ResponseWriter, r *http.Request){
	product, _:= p.DbReader.GetFirstByCode("DC")
	json.NewEncoder(w).Encode(map[string]uint{"id": product.ID})
}

func (p ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request){
	p.DbWriter.WriteToDb(&models.Product{Code: "DC", Price: 100})
	w.WriteHeader(201)
}