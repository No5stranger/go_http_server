package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("service request...")
}

func manHander(name string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println(name)
		fmt.Println("manner service request...")
	})
}

func main() {
	fmt.Println("service start...")
	//http.HandleFunc("/", myHandler)
	//http.ListenAndServe(":7000", nil)
	server := manners.NewWithServer(&http.Server{Addr: ":7000", Handler: manHander("cjp")})
	server.ListenAndServe()
}
