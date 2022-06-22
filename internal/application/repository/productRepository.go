package repository

import "bootcamp/internal/application/models"

// Get spesific product by product Code
func GetProducts() map[string]models.Product {
	products := *models.Storage()
	return products
}

// Get spesific product by product Code
func GetProductById(productCode string) models.Product {
	products := *models.Storage()
	return products[productCode]
}

// Insert new product or update existing
func InsertOrUpdateProduct(payload models.Product) models.Product {
	product := *models.Storage()
	product[payload.Code] = models.Product{
		Code:  payload.Code,
		Name:  payload.Name,
		Price: payload.Price,
	}
	return payload
}

// Delete product by product Code
func DeleteProduct(productCode string) map[string]models.Product {
	product := *models.Storage()
	delete(product, productCode)
	return GetProducts()
}
