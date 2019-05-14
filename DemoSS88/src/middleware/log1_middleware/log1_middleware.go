package log1_middleware

import (
	"fmt"
	"net/http"
)

func Log1Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("url: ", request.URL)
		handler.ServeHTTP(response, request)
	})
}
