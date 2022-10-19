package main

import (
	"fmt"
	"net/http"
)

func getGreeting() string {
	return "Hello, security minicamp!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getGreeting())
}

func main() {
	fmt.Println("Start!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
