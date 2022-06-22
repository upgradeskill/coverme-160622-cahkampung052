package handlers

import (
	"bootcamp/internal/application/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateProduct(t *testing.T) {
	models.GenerateDataDummy()
}
func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actionGet)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetProductById(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/001", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actionGet)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetProductByIdNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/008", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actionGet)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/products/001", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actionDelete)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteHandlerError(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/products/009", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actionDelete)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// func TestPutHandlerError(t *testing.T) {
// 	req, err := http.NewRequest("PUT", "/products/009", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(actionPut)

// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusBadRequest)
// 	}
// }
