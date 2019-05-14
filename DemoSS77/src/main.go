package main

import (
	"apis/demo"
	"apis/mobileapi"
	"apis/productapi"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/demo/demo1", demo.Demo1).Methods("GET")
	router.HandleFunc("/api/demo/demo2", demo.Demo2).Methods("GET")
	router.HandleFunc("/api/demo/demo3/{fullName}", demo.Demo3).Methods("GET")
	router.HandleFunc("/api/product/find", productapi.Find).Methods("GET")
	router.HandleFunc("/api/product/findAll", productapi.FindAll).Methods("GET")
	router.HandleFunc("/api/product/create", productapi.Create).Methods("POST")
	router.HandleFunc("/api/product/update", productapi.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}", productapi.Delete).Methods("DELETE")
	router.HandleFunc("/api/mobile/findAll", mobileapi.FindAll).Methods("GET")

	router.HandleFunc("/api/mobile/search/{keyword}", mobileapi.Search).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
