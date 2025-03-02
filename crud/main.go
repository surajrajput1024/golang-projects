package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	fmt.Printf("Server is up and running on port 9002!")
}

func mainHaindler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		getHandler(w, r)
		break
	case "POST":
		postHandler(w, r)
		break
	case "PUT":
		putHandler(w, r)
		break
	case "DELETE":
		deleteHandler(w, r)
		break
	default:
		{
			http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
			return
		}
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", healthHandler)

	fmt.Printf("Starting server at port 9000\n")
	http.ListenAndServe("localhost:9002", nil)
}
