package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect config
	Conf := config()

	// Add functional
	http.HandleFunc("/", HelloWorld)

	// Listening port
	log.Printf("Starting server at port %s\n", Conf.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", Conf.port), nil); err != nil {
		log.Fatal(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
