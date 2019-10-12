package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var Recovery = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
