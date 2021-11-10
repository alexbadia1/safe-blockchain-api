package main

import (
	"log"
	"net/http"
	"os"
)

//================================================================================
// All User Blockchains stored in a Map.
//
// TODO: Figure out where the heck to put this.
//================================================================================

var UserChains = make(map[string]Blockchain)

//================================================================================
// Documentation Routes
//================================================================================

func docIndex(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/index.html")
} // docIndex

func docCreate(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/create.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/create.html")
} // docCreate

func docChain(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/chain.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/chain.html")
} // docChain

func docMine(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/mine.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/mine.html")
} // docMine

func docImage(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/".
	if r.URL.Path != "/image.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	http.ServeFile(w, r, "templates/image.html")
} // docImage
//================================================================================
// Server Entry Point
//================================================================================

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
	http.HandleFunc("/image.html", docImage)

	// API Endpoints
	http.HandleFunc("/create", new_block)
	http.HandleFunc("/chain", getChain)

	// Images
	http.HandleFunc("/image", getImage)
	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} // main
