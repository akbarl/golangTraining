package demo_api

import (
	"fmt"
	"net/http"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello World")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello")
}

func Demo3(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "World")
}