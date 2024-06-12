package VAS_util_go

import (
	error_handling "github.com/aasulvva/VAS-error-handling-go"
	"net/http"
)

func MethodCheckHandler(next http.HandlerFunc, allowedMethods []string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := false
		for _, method := range allowedMethods {
			if r.Method == method {
				allowed = true
			}
		}

		if !allowed {
			error_handling.LogError(w, error_handling.UnsupportedMethodError(r.Method, allowedMethods))
			return
		}

		next.ServeHTTP(w, r)
	})
}
