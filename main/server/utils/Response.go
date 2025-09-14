package utils

import (
	"context"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

func SetResponseError(r *http.Request, err structs.ApplicationError) {
	*r = *r.WithContext(context.WithValue(r.Context(), "error", err))
}

func GetResponseError(r *http.Request) (structs.ApplicationError, bool) {
	err, ok := r.Context().Value("error").(structs.ApplicationError)
	return err, ok
}

func SetResponseResult[T any](r *http.Request, result T) {
	*r = *r.WithContext(context.WithValue(r.Context(), "result", result))
}

func GetResponseResult[T any](r *http.Request) (T, bool) {
	result, ok := r.Context().Value("result").(T)
	return result, ok
}
