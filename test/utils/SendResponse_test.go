package utils

import (
	"net/http/httptest"
	"testing"
	"zhigalov_tutor_server_core/main/server/utils"

	"github.com/stretchr/testify/assert"
)

func TestSendResponseEmptyJsonArray(t *testing.T) {
	w := httptest.NewRecorder()

	expected := "{\"code\":200,\"message\":\"Success\",\"data\":[]}"
	utils.SendResponse(w, "Success", 200, &[]string{})

	assert.JSONEq(t, expected, w.Body.String())
}

func TestSendResponseJsonObject(t *testing.T) {
	w := httptest.NewRecorder()

	expected := "{\"code\":200,\"message\":\"Success\",\"data\":{\"foo\":\"bar\"}}"
	utils.SendResponse(w, "Success", 200, &struct {
		Foo string `json:"foo"`
	}{Foo: "bar"})

	assert.JSONEq(t, expected, w.Body.String())
}
