package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	// Make sure this endpoint is only accessible at "/index.html".
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/index.html")
} // docIndex

func docCreate(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/create.html".
	if r.URL.Path != "/create.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/create.html")
} // docCreate

func docChain(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/chain.html".
	if r.URL.Path != "/chain.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/chain.html")
} // docChain

func docDelete(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/delete.html".
	if r.URL.Path != "/delete.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/delete.html")
} // docDelete

func docUpdate(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/update.html".
	if r.URL.Path != "/update.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/update.html")
} // docUpdate

func docMine(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/mine.html".
	if r.URL.Path != "/mine.html" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if
	http.ServeFile(w, r, "templates/mine.html")
} // docDelete

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
	router := mux.NewRouter()
	router.HandleFunc("/", docIndex)
	router.HandleFunc("/index.html", docIndex)
	router.HandleFunc("/create.html", docCreate)
	router.HandleFunc("/chain.html", docChain)
	router.HandleFunc("/update.html", docUpdate)
	router.HandleFunc("/delete.html", docDelete)
	router.HandleFunc("/mine.html", docMine)
	router.HandleFunc("/image.html", docImage)

	// API Endpoints
	router.HandleFunc("/create", new_block)
	router.HandleFunc("/chain", getChain)
	router.HandleFunc("/update", update_block)
	router.HandleFunc("/delete", delete_block)

	// Images
	router.HandleFunc("/images", getImage)

	// Wrapping CORS handler
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
} // main
