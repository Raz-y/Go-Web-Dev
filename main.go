package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome</h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact</h1>")
}

func main() {
	var router Router
	fmt.Println("Starting server")

	http.ListenAndServe(":3000", router)
}
