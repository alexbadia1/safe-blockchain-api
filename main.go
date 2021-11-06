package main

import (
	"io"
	"log"
	"net/http"
)

// Map to store all users blockchains
var UserChains = make(map[string]Blockchain)

// Returns the current status of the server
func status(w http.ResponseWriter, r *http.Request) {
	log.Println("\"alive\":true")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
} // status

func main() {
	// TODO: Run locally
	port := "8000"

	// TODO: For production
	// port := os.Getenv("PORT")
	http.HandleFunc("/", status)
	http.HandleFunc("/create", new_block)
	http.HandleFunc("/chain", getChain)
	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} // main
