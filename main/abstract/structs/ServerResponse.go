package structs

import (
	"encoding/json"
	"log"
	"zhigalov_tutor_server_core/main/validation"
)

type ServerResponse[T any] struct {
	Code    uint16 `json:"code" validate:"required,gte=100,lte=599"`
	Message string `json:"message" validate:"required"`
	Data    *T     `json:"data"`
}

func NewServerResponse[T any](code uint16, message string, data *T) (*ServerResponse[T], error) {
	r := &ServerResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

	err := validation.DefaultValidator.Struct(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *ServerResponse[T]) Marshall() *[]byte {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}

	return &jsonResponse
}
