package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server-example start on port 5050")
	http.HandleFunc("/", main_page)

	log.Fatal(http.ListenAndServe(":5050", nil))

}

func main_page(w http.ResponseWriter, r *http.Request) {
	log.Println("Твоя мать шлюха")
	w.Write([]byte("Твоя мать шлюха"))
}
