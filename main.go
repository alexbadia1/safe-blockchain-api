package main

import (
	"log"
	"net/http"
	"os"
)

// Map to store all users blockchains
//
// TODO: Figure out where the heck to put this.
var UserChains = make(map[string]Blockchain)

// Returns the current status of the server
func docIndex(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/index.html")
} // docIndex

// Returns the current status of the server
func docCreate(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/create.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/create.html")
} // docIndex

// Returns the current status of the server
func docChain(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/chain.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/chain.html")
} // docIndex

// Returns the current status of the server
func docMine(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/mine.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/mine.html")
} // docIndex

func main() {
	// TODO: Run locally
	// port := "8000"

	// TODO: For production
	port := os.Getenv("PORT")

	// Documentation Endpoints
	http.HandleFunc("/", docIndex)
	http.HandleFunc("/index.html", docIndex)
	http.HandleFunc("/create.html", docCreate)
	http.HandleFunc("/chain.html", docChain)
	http.HandleFunc("/mine.html", docMine)

	// API Endpoints
	http.HandleFunc("/create", new_block)
	http.HandleFunc("/chain", getChain)
	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} // main
