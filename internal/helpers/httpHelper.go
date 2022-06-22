package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpRespModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Memberikan response api dengan status sukses (HTTP code 200)
func ResponseSuccess(w http.ResponseWriter, message string, data interface{}) {
	resp, _ := json.Marshal(httpRespModel{Message: message, Data: data})

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resp)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}

// Memberikan response api dengan status bad request (HTTP code 400)
func ResponseError(w http.ResponseWriter, message string, data interface{}) {
	resp, _ := json.Marshal(httpRespModel{Message: message, Data: data})

	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write(resp)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}
