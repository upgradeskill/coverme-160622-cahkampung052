package handlers

import (
	"bootcamp/internal/application/models"
	"bootcamp/internal/application/services"
	"bootcamp/internal/helpers"
	"encoding/json"
	"net/http"
	"strings"
)

// Ambil produk code dari array ketiga path url
func getProductCode(r *http.Request) (string, error) {
	productCode := strings.Split(r.URL.Path, "/")
	if len(productCode) >= 3 && productCode[2] != "" {
		return productCode[2], nil
	}

	return "", nil
}

// Decode payload dari json pada raw body http request
func getPayload(w http.ResponseWriter, r *http.Request) models.Product {
	var productPayload models.Product

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&productPayload)

	// Validasi jika format json yang dikirim tidak sesuai dengan model Product
	if err != nil {
		helpers.ResponseError(w, err.Error(), err.Error())
	}

	return productPayload
}

// Mapping routes berdasarkan HTTP Method
func ProductRouting(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		actionGet(w, r)
	case "POST":
		actionPost(w, r)
	case "PUT":
		actionPut(w, r)
	case "DELETE":
		actionDelete(w, r)
	}
}

// Get product
func actionGet(w http.ResponseWriter, r *http.Request) {
	productCode, _ := getProductCode(r)

	// Get all product
	if productCode == "" {
		products := services.GetProducts()
		helpers.ResponseSuccess(w, "success get all product", products)
		return
	}

	// Get spesific product
	product, err := services.GetProductById(productCode)
	if err != nil {
		helpers.ResponseError(w, err.Error(), err)
		return
	}

	helpers.ResponseSuccess(w, "success get single product", product)
}

// Create product
func actionPost(w http.ResponseWriter, r *http.Request) {
	payload := getPayload(w, r)

	// Save new product in memory
	_, err := services.InsertProduct(payload)
	if err != nil {
		helpers.ResponseError(w, err.Error(), err)
		return
	}

	helpers.ResponseSuccess(w, "success create new product", payload)
}

// Update product
func actionPut(w http.ResponseWriter, r *http.Request) {
	payload := getPayload(w, r)

	// Update existing product in memory
	_, err := services.UpdateProduct(payload)
	if err != nil {
		helpers.ResponseError(w, err.Error(), err)
		return
	}

	helpers.ResponseSuccess(w, "success to update product", payload)
}

// Delete product
func actionDelete(w http.ResponseWriter, r *http.Request) {
	productCode, _ := getProductCode(r)

	// Delete product from memory
	err := services.DeleteProduct(productCode)
	if err != nil {
		helpers.ResponseError(w, err.Error(), err)
		return
	}

	helpers.ResponseSuccess(w, "success to delete product", productCode)
}
