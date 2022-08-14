package repository

import (
	"gorm.io/gorm"
	"kube-go-api-tutorial/db/models"
)

type ProductRepository struct{
	Db *gorm.DB
}

func (p ProductRepository) GetFirstByCode(code string) (models.Product, error){
	var product models.Product
	p.Db.First(&product, "code = ?", code)
	return product, nil
}

func (p ProductRepository) WriteToDb(obj *models.Product){
	p.Db.Create(obj)
}
