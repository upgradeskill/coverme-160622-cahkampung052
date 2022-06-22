package services

import (
	"bootcamp/internal/application/models"
	"bootcamp/internal/application/repository"
	"errors"
	"fmt"
)

func GetProducts() []models.Product {
	products := repository.GetProducts()
	return copyToArrayOfMap(products)
}

func GetProductById(productCode string) (p models.Product, err error) {
	product := repository.GetProductById(productCode)
	if len(product.Code) == 0 {
		return p, errors.New("product not found")
	}

	return product, nil
}

func InsertProduct(payload models.Product) (p models.Product, err error) {
	// Validasi product code tidak boleh terdaftar sebelumnya
	_, err = GetProductById(payload.Code)
	if err == nil {
		return p, errors.New("product code already exist")
	}

	// Simpan produk baru ke memory
	repository.InsertOrUpdateProduct(payload)
	return payload, nil
}

func UpdateProduct(payload models.Product) (p models.Product, err error) {
	// Validasi product code sudah terdaftar sebelumnya
	_, err = GetProductById(payload.Code)
	if err != nil {
		fmt.Println(err.Error())
		return p, errors.New("product code is not in our database")
	}

	// Update produk pada memory
	repository.InsertOrUpdateProduct(payload)
	return payload, nil
}

func DeleteProduct(productCode string) error {
	// Validasi product code sudah terdaftar sebelumnya
	_, err := GetProductById(productCode)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("product code is not in our database")
	}

	// Simpan produk baru ke memory
	repository.DeleteProduct(productCode)
	return nil
}

// convert maps menjadi array of object
func copyToArrayOfMap(products map[string]models.Product) []models.Product {
	var newProducts = []models.Product{}
	for _, val := range products {
		newProducts = append(newProducts, val)
	}

	return newProducts
}
