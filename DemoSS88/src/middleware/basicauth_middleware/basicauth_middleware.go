package basicauth_middleware

import (
	"fmt"
	"net/http"
)

func BasicAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		fmt.Println("ok: ", ok)
		if !ok || !checkAccount(username, password) {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
			return
		}
		handler.ServeHTTP(response, request)
	})
}

func checkAccount(username, password string) bool {
	return username == "abc" && password == "123"
}
