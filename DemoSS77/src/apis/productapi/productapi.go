package productapi

import (
	"encoding/json"
	"entities"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func responWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func Find(response http.ResponseWriter, request *http.Request) {
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 11,
	}

	responWithJson(response, http.StatusOK, product)
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	products := []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "name 1",
			Price:    4.5,
			Quantity: 11,
		},
		entities.Product{
			Id:       "p02",
			Name:     "name 2",
			Price:    4.5,
			Quantity: 11,
		},
		entities.Product{
			Id:       "p03",
			Name:     "name 3",
			Price:    4.5,
			Quantity: 11,
		},
	}

	responWithJson(response, http.StatusOK, products)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create new product")
		fmt.Printf("id: %s\nname: %s\nprice: %0.1f\nquantity: %d",
			product.Id, product.Name, product.Price, product.Quantity)
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Update product")
		fmt.Printf("id: %s\nname: %s\nprice: %0.1f\nquantity: %d",
			product.Id, product.Name, product.Price, product.Quantity)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("deleted id: " + id)
}
