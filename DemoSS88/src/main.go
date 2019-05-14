package main

import (
	"demo_api"
	"fmt"
	"middleware/log1_middleware"
	"middleware/log2_middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/api/demo/demo1", log1_middleware.Log1Middleware(http.HandlerFunc(demo_api.Demo1))).Methods("GET")
	router.Handle("/api/demo/demo2", log2_middleware.Log2Middleware(http.HandlerFunc(demo_api.Demo2))).Methods("GET")
	router.Handle("/api/demo/demo3", log2_middleware.Log2Middleware(http.HandlerFunc(demo_api.Demo3))).Methods("GET")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
	}
}
