package jwt_middleware

import (
	"encoding/json"
	"net/http"
	"security"

	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		stringToken := request.Header.Get("token")
		if stringToken == "" {
			responseWithError(response, http.StatusUnauthorized, "Unauthorized")
		} else {
			result, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(security.PrivateKey), nil
			})
			if err != nil {
				responseWithError(response, http.StatusUnauthorized, err.Error())
			} else {
				if result.Valid {
					handler.ServeHTTP(response, request)
				} else {
					responseWithError(response, http.StatusUnauthorized, "Invalid key")
				}
			}
		}
	})
}

func responWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responWithJson(response, statusCode, map[string]string{"error": msg})
}
