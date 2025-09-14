package middlewares

import (
	"encoding/json"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/structs"
	"zhigalov_tutor_server_core/main/server/utils"
)

func ResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)

		if err, ok := utils.GetResponseError(r); ok {
			type errorBody struct {
				Message string `json:"text"`
			}

			response, e := structs.NewServerResponse[errorBody](err.Code, err.Message, &errorBody{Message: err.Error()})
			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			e = encoder.Encode(response)
			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		}

		result, ok := utils.GetResponseResult[any](r)
		if !ok {
			result = struct{}{}
		}

		success, err := structs.NewServerResponse[any](http.StatusOK, "success", &result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.(http.Flusher).Flush()
			return
		}

		err = encoder.Encode(success)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.(http.Flusher).Flush()
			return
		}

		w.(http.Flusher).Flush()
		return
	})
}
