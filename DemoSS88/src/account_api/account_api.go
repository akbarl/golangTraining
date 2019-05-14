package account_api

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"

	"gopkg.in/mgo.v2/bson"
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

func Create(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		responseWithError(response, http.StatusBadGateway, err.Error())
	} else {
		db, err2 := config.GetConnect()
		if err2 != nil {
			responseWithError(response, http.StatusBadGateway, err2.Error())
		} else {
			accountModel := models.AccountModel{
				Db:         db,
				Collection: "account",
			}

			account.Id = bson.NewObjectId()
			err3 := accountModel.Create(&account)
			if err3 != nil {
				responseWithError(response, http.StatusBadGateway, err2.Error())
			} else {
				responWithJson(response, http.StatusOK, account)
			}
		}
	}
}
