package services

import (
	"bootcamp/internal/application/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateProduct(t *testing.T) {
	models.GenerateDataDummy()
}

func TestGetAllProduct(t *testing.T) {
	product := GetProducts()
	assert := assert.New(t)
	assert.Equal(len(product), 3, "Total Seharusnya 3")
}

func TestGetAProductById(t *testing.T) {
	product, _ := GetProductById("002")
	assert := assert.New(t)
	assert.Equal(product.Code, "002", "Total Seharusnya 3")
}

func TestInsertProductSuccess(t *testing.T) {
	product, _ := InsertProduct(models.Product{
		Code:  "005",
		Name:  "Susu Murni",
		Price: 1000,
	})

	assert.Equal(t, product.Code, "005", "Code seharusnya 005")
}

func TestInsertProductError(t *testing.T) {
	_, err := InsertProduct(models.Product{
		Code:  "002",
		Name:  "Susu Murni",
		Price: 1000,
	})

	assert.Error(t, err, "seharusnya error, karena kode sudah terdaftar")
}

func TestUpdateProductSuccess(t *testing.T) {
	product, _ := UpdateProduct(models.Product{
		Code:  "002",
		Name:  "Susu Murni",
		Price: 1000,
	})

	assert.Equal(t, product.Code, "002", "Code seharusnya 002")
}

func TestUpdateProductError(t *testing.T) {
	_, err := UpdateProduct(models.Product{
		Code:  "008",
		Name:  "Susu Murni",
		Price: 1000,
	})

	assert.Error(t, err, "seharusnya error karena kode tidak terdaftar")
}

func TestDeleteProductSuccess(t *testing.T) {
	err := DeleteProduct("005")
	assert.Equal(t, err, nil, "Seharusnya tidak error")
}

func TestDeleteProductError(t *testing.T) {
	err := DeleteProduct("005")
	assert.Error(t, err, "seharusnya error karena kode sudah dihapus pada test sebelumnya")
}
