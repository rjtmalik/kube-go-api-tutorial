package repository_interfaces

import "kube-go-api-tutorial/db/models"

type ProductReader interface{
	GetFirstByCode(code string) (models.Product, error)
}

type ProductWriter interface{
	WriteToDb(*models.Product)
}