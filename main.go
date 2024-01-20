package main

import (
	"fmt"
	"log"

	"net/http"
)

func dummyMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("####### YO I am the middlware #######")
		next.ServeHTTP(w, req)
	})
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("[GET] [/hello]")
	fmt.Fprintf(w, "Hello\n")

}

func meow(w http.ResponseWriter, req *http.Request) {
	log.Println("[GET] [/meow]")
	fmt.Fprintf(w, "Meow\n")
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func main() {
	meowHandler := http.HandlerFunc(meow)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/meow", dummyMiddleware(meowHandler))

	port := 5050
	log.Printf("Server is listening on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	checkError(err, "Error starting server")
}
