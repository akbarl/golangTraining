package log2_middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Log2Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		today := time.Now()
		fmt.Println("Date: " + today.Format("01/02/2006 15:04:05"))
		handler.ServeHTTP(response, request)
	})
}
