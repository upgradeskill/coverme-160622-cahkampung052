package models

type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductStorage struct{}

// Make data map to save product in memory
var products = make(map[string]Product)

// Generate data dummy and save to memory
func GenerateDataDummy() {
	products["001"] = Product{Code: "001", Name: "Susu Sapi", Price: 10000}
	products["002"] = Product{Code: "002", Name: "Susu Kambing", Price: 10000}
	products["003"] = Product{Code: "003", Name: "Susu Kuda", Price: 10000}
}

// Get data product from memory
func Storage() *map[string]Product {
	return &products
}
