package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendResponse(t *testing.T) {
	type args[T any] struct {
		w       http.ResponseWriter
		message string
		code    uint16
		data    *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[struct{}]{
		{
			name: "Send OK response",
			args: args[struct{}]{
				w:       httptest.NewRecorder(),
				code:    http.StatusOK,
				message: "Success",
				data:    new(struct{}),
			},
		},
		{
			name: "Send NotFound response",
			args: args[struct{}]{
				w:       httptest.NewRecorder(),
				code:    http.StatusNotFound,
				message: "Not Found",
				data:    new(struct{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendResponse(tt.args.w, tt.args.message, tt.args.code, tt.args.data)
		})
	}

	t.Run("Send response with function as data", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SendResponse() did not panic")
			}
		}()

		f := func() {}

		SendResponse(httptest.NewRecorder(), "I'm should panic", http.StatusInternalServerError, &f)
	})

	t.Run("Send response with not valid data", func(t *testing.T) {
		type NotValidData struct {
			Text string `json:"data" validate:"required"`
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SendResponse() did not panic")
			}
		}()

		SendResponse(httptest.NewRecorder(), "I'm should panic", http.StatusBadRequest, &NotValidData{})
	})
}
