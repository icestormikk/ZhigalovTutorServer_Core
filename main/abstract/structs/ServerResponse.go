package structs

import (
	"encoding/json"
	"log"
)

type ServerResponse struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *ServerResponse) Marshall() *[]byte {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}

	return &jsonResponse
}
