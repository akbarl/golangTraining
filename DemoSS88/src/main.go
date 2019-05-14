package main

import (
	"account_api"
	"demo_api"
	"fmt"
	"jwt_api"
	"middleware/basicauth_middleware"
	"middleware/jwt_middleware"
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
	router.Handle("/api/demo/demo4", basicauth_middleware.BasicAuthMiddleware(http.HandlerFunc(demo_api.Demo4))).Methods("GET")

	router.HandleFunc("/api/account/create", account_api.Create).Methods("POST")
	router.HandleFunc("/api/generate/token", jwt_api.GenerateToken).Methods("POST")
	router.Handle("/api/demo/demo5", jwt_middleware.JWTMiddleware(http.HandlerFunc(demo_api.Demo5))).Methods("GET")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
	}
}
