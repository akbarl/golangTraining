package jwt_api

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"
	"security"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func responWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responWithJson(response, statusCode, map[string]string{"error": msg})
}

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		db, err2 := config.GetConnect()
		if err2 != nil {
			responseWithError(response, http.StatusBadGateway, err2.Error())
		} else {
			accountModel := models.AccountModel{
				Db:         db,
				Collection: "account",
			}
			isValid := accountModel.Check(account.Username, account.Password)

			if !isValid {
				responseWithError(response, http.StatusBadRequest, "Invalid")
			} else {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": account.Username,
					"password": account.Password,
					"exp":      time.Now().Add(20 * time.Hour).Unix(),
				})
				tokenString, err2 := token.SignedString([]byte(security.PrivateKey))
				if err2 != nil {
					responseWithError(response, http.StatusBadRequest, err2.Error())
				} else {
					responWithJson(response, http.StatusOK, entities.JWToken{
						Token: tokenString,
					})
				}
			}
		}
	}
}
