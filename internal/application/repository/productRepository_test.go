package repository

import (
	"bootcamp/internal/application/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	products := GetProducts()
	assert.Equal(t, len(products), 0, "Total seharusnya 0")
}

func TestGetProductsById(t *testing.T) {
	products := GetProductById("002")
	assert.Equal(t, products.Code, "", "Code seharusnya kosong karena tidak ada data")
}

func TestInsertOrUpdate(t *testing.T) {
	products := InsertOrUpdateProduct(models.Product{
		Code:  "008",
		Name:  "Susu Murni",
		Price: 1000,
	})
	assert.Equal(t, products.Code, "008", "Code sesuai dengan payload yang dikirim")

	newProduct := GetProducts()
	assert.Equal(t, len(newProduct), 1, "Total produk seharusnya 1 karena baru saja diinput")
}

func TestDelete(t *testing.T) {
	products := DeleteProduct("008")
	assert.Equal(t, len(products), 0, "Produk seharusnya 0 karena baru saja dihapus")
}
