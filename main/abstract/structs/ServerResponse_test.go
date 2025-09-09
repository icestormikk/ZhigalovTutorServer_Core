package structs

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewServerResponse(t *testing.T) {
	type args[T any] struct {
		code    uint16
		message string
		data    *T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    *ServerResponse[T]
		wantErr bool
	}

	tests := []testCase[struct{}]{
		{
			name:    "Successful response",
			args:    args[struct{}]{code: http.StatusOK, message: "Success", data: &struct{}{}},
			want:    &ServerResponse[struct{}]{Code: http.StatusOK, Message: "Success", Data: &struct{}{}},
			wantErr: false,
		},
		{
			name:    "Failed response",
			args:    args[struct{}]{code: http.StatusNotFound, message: "Not found", data: &struct{}{}},
			want:    &ServerResponse[struct{}]{Code: http.StatusNotFound, Message: "Not found", Data: &struct{}{}},
			wantErr: false,
		},
		{
			name:    "Response with incorrect http code",
			args:    args[struct{}]{code: 0, message: "Incorrect http code", data: &struct{}{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Response with empty message",
			args:    args[struct{}]{code: http.StatusOK, message: "", data: &struct{}{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServerResponse(tt.args.code, tt.args.message, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServerResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServerResponse() got = %v, want %v", got, tt.want)
			}
		})
	}

	type node struct {
		Val  bool  `json:"val"`
		Next *node `json:"next"`
	}
	n1 := node{Val: true}
	n2 := node{Val: true, Next: &n1}
	n1.Next = &n2

	t.Run("Response with circular data", func(t *testing.T) {
		if recover() != nil {
			_, err := NewServerResponse(http.StatusInternalServerError, "Circular structure", &n1)
			if err == nil {
				t.Errorf("NewServerResponse() error expected, but not received")
				return
			}
		}
	})
}

func TestServerResponse_Marshall(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    *ServerResponse[T]
		want []byte
	}

	tests := []testCase[struct{}]{
		{
			name: "Response json with correct data",
			r:    &ServerResponse[struct{}]{Code: http.StatusOK, Message: "Success", Data: new(struct{})},
			want: []byte("{\"code\":200,\"message\":\"Success\",\"data\":{}}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Marshall(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Marshall() = %v, want %v", *got, tt.want)
			}
		})
	}

	type Node struct {
		Val  bool  `json:"val"`
		Next *Node `json:"next"`
	}
	n1 := &Node{Val: true}
	n2 := &Node{Val: true, Next: n1}
	n1.Next = n2

	r := &ServerResponse[Node]{Code: http.StatusOK, Message: "Success but with surprise", Data: n1}

	t.Run("Response json with incorrect data", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("NewServerResponse() error expected, but not received")
			}
		}()

		r.Marshall()
	})
}
