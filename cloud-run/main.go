package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := "World"
	if r.URL.Path != "/" {
		name = r.URL.Path[1:]
	}
	msg := fmt.Sprintf("Hello, %s!\n", name)
	fmt.Fprint(w, msg)
}

func main() {
	port := ":8080"
	log.Printf("listening on %s\n", port)
	http.ListenAndServe(":8080", http.HandlerFunc(Handler))
}
