package internal

import (
	"bootcamp/internal/application/handlers"
	"bootcamp/internal/application/models"
	"fmt"
	"net/http"
)

func Main() {
	// Generate data dummy
	models.GenerateDataDummy()

	// Handler request
	http.HandleFunc("/products/", handlers.ProductRouting)
	fmt.Println("starting web server at http://localhost:8000/")

	// Start server
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
