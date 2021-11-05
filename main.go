package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func status(w http.ResponseWriter, r *http.Request) {
	log.Println("\"alive\":true")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
} // status

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", status)
	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} // main
