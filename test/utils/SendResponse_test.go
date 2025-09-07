package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"zhigalov_tutor_server_core/main/server/utils"

	"github.com/stretchr/testify/assert"
)

func TestSendResponseEmptyJsonArray(t *testing.T) {
	w := httptest.NewRecorder()

	expected := "{\"code\":200,\"message\":\"Success\",\"data\":[]}"
	utils.SendResponse(w, "Success", http.StatusOK, &[]string{})

	assert.JSONEq(t, expected, w.Body.String())
}

func TestSendResponseJsonObject(t *testing.T) {
	w := httptest.NewRecorder()

	expected := "{\"code\":200,\"message\":\"Success\",\"data\":{\"foo\":\"bar\"}}"
	utils.SendResponse(w, "Success", http.StatusOK, &struct {
		Foo string `json:"foo"`
	}{Foo: "bar"})

	assert.JSONEq(t, expected, w.Body.String())
}

func TestSendResponseError(t *testing.T) {
	w := httptest.NewRecorder()

	type Node struct {
		Value int   `json:"value" validate:"-"`
		Next  *Node `json:"next" validate:"-"`
	}

	n1 := &Node{Value: 0}
	n2 := &Node{Value: 1, Next: n1}
	n1.Next = n2

	assert.Panics(t, func() {
		utils.SendResponse[Node](w, "", http.StatusInternalServerError, n1)
	})
}
