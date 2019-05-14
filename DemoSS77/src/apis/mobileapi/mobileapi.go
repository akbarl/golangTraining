package mobileapi

import (
	"config"
	"encoding/json"
	"fmt"
	"models"
	"net/http"

	"github.com/gorilla/mux"
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

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetConnect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db:         db,
			Collection: "mobile",
		}
		mobiles, err2 := mobileModel.FindAll()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			fmt.Println(len(mobiles))
			responWithJson(response, http.StatusOK, mobiles)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetConnect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db:         db,
			Collection: "mobile",
		}
		vars := mux.Vars(request)
		keyword := vars["keyword"]
		mobiles, err2 := mobileModel.Search(keyword)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			fmt.Println(len(mobiles))
			responWithJson(response, http.StatusOK, mobiles)
		}
	}
}
