package structs

import (
	"net/http"
	"testing"
	"zhigalov_tutor_server_core/main/abstract/structs"

	"github.com/stretchr/testify/assert"
)

func TestServerResponseMarshallSuccess(t *testing.T) {
	response, _ := structs.NewServerResponse(
		http.StatusOK,
		"Success",
		&struct {
			Foo string `json:"foo"`
		}{Foo: "bar"},
	)

	expected := "{\"code\":200,\"message\":\"Success\",\"data\":{\"foo\":\"bar\"}}"

	assert.Equal(t, expected, string(*response.Marshall()))
}

func TestServerResponseNoMessage(t *testing.T) {
	_, err := structs.NewServerResponse(
		http.StatusInternalServerError,
		"",
		&struct {
			Foo string `json:"foo"`
		}{Foo: "bar"},
	)

	assert.Error(t, err)
}

func TestServerResponseZeroCode(t *testing.T) {
	_, err := structs.NewServerResponse(
		0,
		"Zero code",
		&struct {
			Foo string `json:"foo"`
		}{Foo: "bar"},
	)

	assert.Error(t, err)
}

func TestServerResponseMarshallBadData(t *testing.T) {
	type Node struct {
		Value bool  `json:"value" validate:"-"`
		Next  *Node `json:"next" validate:"-"`
	}

	n1 := &Node{Value: true}
	n2 := &Node{Value: true, Next: n1}
	n1.Next = n2

	response, err := structs.NewServerResponse(
		http.StatusInternalServerError,
		"Error",
		n1,
	)

	if err != nil {
		t.Error(err)
	}

	assert.Panics(t, func() { response.Marshall() })
}
