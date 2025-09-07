package utils

import (
	"encoding/json"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

func SendResponse[T any](w http.ResponseWriter, message string, code uint16, data *T) {
	encoder := json.NewEncoder(w)

	response := &structs.ServerResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	err := encoder.Encode(response)
	if err != nil {
		panic("Error while marshalling data")
	}
}
